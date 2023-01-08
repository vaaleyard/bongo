

<div align="center">
<h1> bongo </h1>
A mongo management TUI dashboard <br>

<br>

![lines of code](https://sloc.xyz/github/vaaleyard/bongo) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/bongo) ![help wanted](https://img.shields.io/github/labels/vaaleyard/bongo/help%20wanted) ![good first issue](https://img.shields.io/github/labels/vaaleyard/bongo/good%20first%20issue) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

## Screenshot
![screenshot](./assets/screenshot.png)

## Features
- Navigate through database collections, views and users in a tree view
- Run a [database command](https://www.mongodb.com/docs/manual/reference/command/) and view it's output (you can also copy it to the clipboard)
- Vim keybindings

## Installation
### With the binary
Download the binary for your distribution in the releases page.
### With `go install`
```shell
go install github.com/vaaleyard/bongo@latest
```
### With make
```
git clone https://github.com/vaaleyard/bongo.git
make build 
```

## Usage
Export your mongodb uri and run bongo:
```bash
export MONGODB_URI='mongodb://user:password@host:27017/'
bongo --mongodb-uri $MONGODB_URI
```

## Keybindings
```
esc: go to normal mode (unfocus)
d/D: focus the database finder block
i/I/: (colon) : focus the input block
p/P: focus the preview block
y/Y: (if preview is focused) [yank/copy] the content of preview block
s/S: (if database tree is focused) select the database to which database the command in input block will run
q/Q/Ctrl-c: quit
```

## Features to do
- [ ] Make a preview when viewing the items in the tree
- [ ] Make an input field to run commands against the mongo connection, i.g. `db.getUsers()`  
- [ ] Create multiple connections with tabs
- [ ] Improve colorscheme
- [ ] Improve mongo connection handling
- [ ] Create a help page/widget
- [ ] Better cluster handling (maybe one page to add connections, etc.)
- [ ] Expand input box if user paste a big mongo command

## Contributing
See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License
[MIT](./LICENSE)
