import argparse
import json
import os.path
import subprocess
import urllib.request

from hashlib import sha1
from io import BufferedReader
from typing import Any, Dict

NETLIFY_BASE_URL = "https://api.netlify.com/api/v1"


def render_pdf_file(input_file: str) -> None:
    try:
        subprocess.Popen(["xelatex", input_file])
    except FileNotFoundError as e:
        print(e)
        raise SystemExit


def get_file_sha(input_file: str) -> str:
    try:
        digest: str = None
        with open(input_file, "rb") as f:
            contents = f.read()
            digest = sha1(contents).hexdigest()

        return digest
    except FileNotFoundError:
        print(e)
        raise SystemExit


def prepare_file(input_file: str) -> (BufferedReader, str):
    render_pdf_file(input_file)

    # Create new filename and generate the new path.
    new_filename = os.path.basename(input_file).split(".")[0]
    new_file_path = f"{os.path.dirname(input_file)}/{new_filename}.pdf"
    file_contents = open(new_file_path, "rb")

    # Get SHA1SUM.
    sha = get_file_sha(new_file_path)

    return file_contents, sha


def do_netlify_request(
    path: str, headers: Dict[str, Any], method: str = "GET", data=None
) -> Dict[Any, Any]:
    req = urllib.request.Request(
        url=f"{NETLIFY_BASE_URL}{path}", data=data, headers=headers, method=method
    )

    response: Dict[Any, Any] = None
    with urllib.request.urlopen(req) as resp:
        response = json.loads(resp.read())

    return response


def get_most_recent_deploy_meta(site_url: str, auth_header: Dict[str, str]) -> Dict[str, str]:
    resp = do_netlify_request(f"/sites/{site_url}/deploys", headers=auth_header)

    if not resp:
        print(f"Site {site_url} not found!")
        raise SystemExit

    return {"id": resp[0]["id"], "updated_at": resp[0]["updated_at"], "state": resp[0]["state"]}


def get_deploy_file_meta(deploy_id: str, auth_header: Dict[str, str]) -> Dict[str, Dict[str, str]]:
    resp = do_netlify_request(f"/deploys/{deploy_id}/files", headers=auth_header)

    if not resp:
        print("Could not get list of most recent file. Exiting...")
        raise SystemExit

    files = {"files": {}}
    for file in resp:
        name = file["id"]
        sha = file["sha"]

        files["file"][name] = sha

    return files


def upload_rendered_file(
    url: str, token: str, file_path: str, sha: str, contents: BufferedReader
) -> None:
    auth_header = {"Authorization": f"Bearer {token}"}

    # Get previous deployment information.
    deploy_meta = get_most_recent_deploy_meta(url, auth_header)
    file_meta = get_deploy_file_meta(deploy_meta["id"], auth_header)

    if file_path in file_meta["files"]:
        if file_meta["files"][file_path] == sha:
            return

    # Update or insert SHA1SUM of file.
    file_meta["files"][file_path] = sha

    # Initiate new deployment and get its ID.
    do_netlify_request(f"/sites/{url}/deploys/", headers=auth_header, method="POST", data=file_meta)
    new_meta = get_most_recent_deploy_meta(url, auth_header)

    # Upload files and finalize deployment.
    upload_headers = auth_header.copy()
    upload_headers["Content-Type"] = "application/octet-stream"

    do_netlify_request(
        f"/deploys/{new_meta['id']}/files/{file_path}",
        headers=upload_headers,
        method="PUT",
        data=contents,
    )


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--site-url", type=str, dest="url", require=True)
    parser.add_argument("--input-file", type=str, dest="file", required=True)
    parser.add_argument("--token", type=str, required=True)
    parser.add_argument(
        "--web-path",
        type=str,
        help="Path (not including filename) of file in CDN.",
        dest="wpath",
        required=True,
    )
    args = parser.parse_args()

    # Buffered contents of file and SHA1SUM of file.
    contents, sha = prepare_file(args.file)

    # Create new deployment and upload to Netlify.
    web_file_path = f"{args.wpath}{contents.name}"
    upload_rendered_file(args.url, args.token, web_file_path, sha, contents)


if __name__ == "__main__":
    main()
