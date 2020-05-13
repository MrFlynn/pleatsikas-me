# pleatsikas-me

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