# Contributing
All contributions are welcome and I will be extremely grateful if you make it.
Since this project is in it's very early stages, there are a lot of things you can contribute with. You can start out by reading the project [page](https://github.com/users/vaaleyard/projects/1/views/1). Go through all the items and choose one of them to contribute.

## Running
First, you will need a running mongo database. I've made a docker-compose to make our life easier:
```
cd src/mongo
docker compose up -d
export MONGODB_URI='mongodb://admin:bergo@localhost:27017/'
```
Then, you can execute the project with cargo basic commands: `cargo run`

## Pull Requests
1. Fork the project.
2. Make your changes locally.
3. Commit and create a Pull Request. In the text, for now, just explain what you did and that's it. There's not a template, yet.
4. After you've created it, just wait a little, I will review it and merge it if it's everything is alright.

## What you need to know to contribute
Just to clarify, I am learning rust as this project evolves. So, of course I will also not know things and I will also have to study a lot of stuff to finish this project (probably you will know rust more than me). But, if you want to contribute, you can start out by learning the [tui-rs](https://github.com/fdehau/tui-rs/) library, because it's used to display all the User Interface and the [crossterm](https://github.com/crossterm-rs/crossterm) library for the input handling (there's an example in the tui-rs repository).
