# bongo
This is a prototype of a program to navigate through mongo resources (databases, collections, etc.)
and run commands to it. Basically, it's almost a k9s version for mongo databases.

<div align="center">

[![made-with-rust](https://img.shields.io/badge/Made%20with-Rust-1f425f.svg)](https://www.rust-lang.org/) ![Code Size](https://img.shields.io/github/languages/code-size/vaaleyard/bongo) ![help wanted](https://img.shields.io/github/labels/vaaleyard/bongo/help%20wanted) ![good first issue](https://img.shields.io/github/labels/vaaleyard/bongo/good%20first%20issue) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

</div> 

# Screenshot
![screenshot](./assets/screenshot.png)

# Features to do
- [ ] Make a database tree (expand each collection to a tree to view collections, views, users, etc.)
and navigate through them.  
- [ ] Make a preview when viewing those items above  
- [ ] Make an input field to run commands against the mongo connection, i.g. `db.getUsers()`  
- [ ] Create multiple connections

## Contributing
See [CONTRIBUTING.md](./CONTRIBUTING.md).

## License
[MIT](./LICENSE)
