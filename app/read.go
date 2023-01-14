package app

import(
	"github.com/edemond/wiki/context"
	"github.com/edemond/wiki/graph"
	"github.com/edemond/wiki/log"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var _readPageTemplate *template.Template

func init() {
	_readPageTemplate = mustParseTemplates("read.html", "read.html", "base.html")
}

type ReadPage struct {
	Alias graph.Alias
	Contents template.HTML
	Title string
}

// readPage displays a page given an alias.
func readPage(rw http.ResponseWriter, r *http.Request) {
	ctx := context.GetHTTPContext(r)

	vars := mux.Vars(r)
	alias := vars["alias"]

	ns := graph.NewNodeSource()

	// TODO: Save and load nodes by the ID, not the alias. The alias
	// should be looked up by a different service, for indirection.
	node, err := ns.GetNode(ctx, graph.ID(alias))
	if err != nil {
		log.Printf(ctx, "Error getting node for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusNotFound) // TODO: hmmmm disambiguate this from error situations
		return
	} else if node == nil {
		// TODO: 404 page.
		log.Printf(ctx, "No node found for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	innerHTML, err := getHTMLFromMarkdown(node.Contents)
	if err != nil {
		log.Errorf(ctx, "Error converting Markdown to HTML for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = _readPageTemplate.ExecuteTemplate(rw, "base", &ReadPage{
		Alias: node.Alias,
		Contents: template.HTML(innerHTML), // TODO: danger will robinson
		Title: node.Title,
	})
	if err != nil {
		log.Errorf(ctx, "Error executing read page template for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}
