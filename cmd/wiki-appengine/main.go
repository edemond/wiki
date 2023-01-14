package main

import (
	"github.com/edemond/wiki/app"
	"google.golang.org/appengine"
	"net/http"

	_ "github.com/edemond/wiki/context/appengine"
	_ "github.com/edemond/wiki/graph/datastore"
	_ "github.com/edemond/wiki/search/appengine"
	_ "github.com/edemond/wiki/log/appengine"
)

func main() {
	// TODO: Are we supposed to register routes in init()?
	// Any problem with doing it here?
	router := app.GetRouter(false) // Don't serve static files; those are set up in app.yaml
	http.Handle("/", router)
	appengine.Main()
}
