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

/* TODO:
 * Keep only one connection in mongo and close the context at the end
 * Improve colors and focus
 * RunCommand against mongo when hitting Enter in the input box
 * Create Help page
 * Better cluster handling (maybe one page to add connections, etc.)
 * Expand input box if user paste a big mongo command
 * Multi Connections with Tabs
 */
