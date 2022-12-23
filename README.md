

<div align="center">
<h1> bongo </h1>
A mongo management TUI dashboard <br>

<br>

![lines of code](https://sloc.xyz/github/vaaleyard/bongo) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/bongo) ![help wanted](https://img.shields.io/github/labels/vaaleyard/bongo/help%20wanted) ![good first issue](https://img.shields.io/github/labels/vaaleyard/bongo/good%20first%20issue) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

## Screenshot
![screenshot](./assets/screenshot.png)

## Installation
### With make
```
git clone https://github.com/vaaleyard/bongo.git
make build 
```

## Usage
Export your mongodb uri and run bongo:
```bash
export MONGODB_URI='mongodb://user:password@host:27017/'
./bongo/bongo --mongodb-uri $MONGODB_URI
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
