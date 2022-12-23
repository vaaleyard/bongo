package main

import (
	"flag"
	"github.com/vaaleyard/bongo/bongo"
)

func main() {
	uri := flag.String("mongodb-uri", "", "URI of the mongoDB you want to connect.")
	flag.Parse()

	app, _ := bongo.Init(*uri)
	bongo.Ui(app)
}
