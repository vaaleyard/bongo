# Contributing
All contributions are welcome and I will be extremely grateful if you make it.
Since this project is in it's very early stages, there are a lot of things you can contribute with. You can start out
by reading the project [page](https://github.com/users/vaaleyard/projects/1/views/1).
Go through all the items and choose one of them to contribute.

## Running
First, you will need a running mongo database. I've made a docker-compose to make our life easier:
```
cd src/mongo
docker compose up -d
export MONGODB_URI='mongodb://admin:admin@localhost:27017/'
```
Then, you can execute the project with the go run command: `go run . --mongodb-uri $MONGODB_URI`

## Pull Requests
1. Fork the project.
2. Make your changes locally.
3. Commit and create a Pull Request. In the text, for now, just explain what you did and that's it. There's not a template, yet.
4. After you've created it, just wait a little, I will review it and merge it if it's everything is alright.

## What you need to know to contribute
If you want to contribute, you can start out by learning the [tview](https://github.com/rivo/tview/) library. There are
a lot of examples in this library showing how to use each widget, running them is a good start. The code itself isn't
hard, so I think it will be easy to understand what each line is doing.
