package main

import "github.com/vaaleyard/bongo/bongo"

func main() {
	app := bongo.Init()
	bongo.Ui(app)
}
