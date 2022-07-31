

<div align="center">
<h1> bongo </h1>
A mongo management TUI dashboard <br>

<br>

![lines of code](https://sloc.xyz/github/vaaleyard/bongo) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/bongo) ![help wanted](https://img.shields.io/github/labels/vaaleyard/bongo/help%20wanted) ![good first issue](https://img.shields.io/github/labels/vaaleyard/bongo/good%20first%20issue) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

## Screenshot
![screenshot](./assets/screenshot.png)

## Installation
### With cargo
```
git clone https://github.com/vaaleyard/bongo.git
cargo install --path bongo
```

## Usage
Export your mongodb uri and run bongo:
```bash
export MONGODB_URI='mongodb://user:password@host:27017/'
bongo
```

## Features to do
- [ ] Make a database tree (expand each collection to a tree to view collections, views, users, etc.)
and navigate through them.  
- [ ] Make a preview when viewing those items above  
- [ ] Make an input field to run commands against the mongo connection, i.g. `db.getUsers()`  
- [ ] Create multiple connections

## Contributing
See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License
[MIT](./LICENSE)
