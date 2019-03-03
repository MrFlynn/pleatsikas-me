import json
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


def do_netlify_request(path: str, headers: Dict[str, Any], method: str = "GET", data=None) -> dict:
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

