/* 
	Package main implements the "plain" version of the wiki,
	which self-hosts with the net/http server.

	It's designed to store all of its data locally. This may change
	as we gain the ability to call out to databases running on
	different hosts. Or, it's conceivable that we could self-host 
	the web app but still use Google Cloud Platform storage, etc.
*/
package main

import (
	"github.com/edemond/wiki/app"
	"github.com/edemond/wiki/graph"
	"flag"
	"fmt"
	"net/http"
	"strings"

	// Dependencies 
	_ "github.com/edemond/wiki/graph/cayley"
	_ "github.com/edemond/wiki/graph/file"
	_ "github.com/edemond/wiki/search/bleve"
)

var(
	_readAlias = flag.String("r", "", "Read a page by the given alias.")
	_writeEdge = flag.String("e", "", "Write a new edge. Format: from edge to") 
)

func main() {
	// TODO: This should be split out into a multitool kind of command 
	flag.Parse()
	if *_readAlias != "" {
		fmt.Println("read mode")
	} else if *_writeEdge != "" {
		err := writeEdge(*_writeEdge)
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else {
		serve()
	}
}

func serve() {
	router := app.GetRouter(true) // Serve static files ourselves.
	fmt.Printf("Now serving on localhost:8080.")
	http.ListenAndServe(":8080", router)
}

func writeEdge(input string) error {
	// Expect a string like: <from> <edge> <to>
	nquad := strings.Split(input, " ")
	if len(nquad) != 3 {
		return fmt.Errorf("Expected format: <from ID> <edge name> <to ID>\n")
	}
	fmt.Printf("Writing edge: from:%v edge:%v to:%v",
		nquad[0], nquad[1], nquad[2],
	)

	// TODO: Check that these actually exist in the graph.
	g := graph.NewGraph()
	err := g.AddEdge(
		graph.ID(nquad[0]),
		nquad[1],
		graph.ID(nquad[2]),
	)
	if err != nil {
		return fmt.Errorf("Couldn't create edge: %v", err)
	}

	return nil
}
