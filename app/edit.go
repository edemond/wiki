package app

import(
	"github.com/edemond/wiki/context"
	"github.com/edemond/wiki/graph"
	"github.com/edemond/wiki/log"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

var _editPageTemplate *template.Template

func init() {
	_editPageTemplate = mustParseTemplates(
		"edit.html",
		"edit.html", 
		"base.html",
	)
}

type EditPage struct {
	ID string
	Alias graph.Alias
	Contents graph.Markdown
	Title string
}

// editPage loads a page in the editor or handles an edit page submission.
func editPage(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		editPageGet(rw, r)
	} else if r.Method == "POST" {
		newOrEditPagePost(rw, r) 
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed) 
	}
}

func editPageGet(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	ns := graph.NewNodeSource()
	ctx := context.GetHTTPContext(r)

	// TODO: Save and load nodes by the ID, not the alias. The alias
	// should be looked up by a different service, for indirection.
	node, err := ns.GetNode(ctx, graph.ID(alias))
	if err != nil {
		log.Errorf(ctx, "Error getting node for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else if node == nil {
		// TODO: 404 page.
		log.Errorf(ctx, "No node found for alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	_editPageTemplate.ExecuteTemplate(rw, "base", &EditPage{
		ID: node.ID,
		Alias: node.Alias,
		Title: node.Title,
		Contents: node.Contents,
	})
}
