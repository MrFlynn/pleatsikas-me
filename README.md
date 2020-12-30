# pleatsikas-me

[![Netlify Status](https://api.netlify.com/api/v1/badges/2bc59369-5924-45d7-984e-fe4c17b81251/deploy-status)](https://app.netlify.com/sites/pleatsikas-me/deploys)

The code for my personal website. It is written in [Vue](https://vuejs.org) and 
compiled to as close to a static page I can get with some dynamic elements 
(i.e. modals and the like).

This site also automatically deploys to Netlify using their build environment
and Github actions to compile and upload [my resume](https://pleatsikas.me/resume.pdf).

## Development
You can use the following commands to develop the website locally:
```bash
$ yarn serve # Create a local webserver with unoptimized build and hot reload.
$ yarn build # Create release build.
```

To build a pdf version of my resume you need Latexmk with XeLaTeX installed.
In order to generate my resume from a template run the following:
```bash
$ cd src/assets
$ go run ../../tools/resume-builder/main.go
```

_Note:_ You will need the Go toolchain installed.