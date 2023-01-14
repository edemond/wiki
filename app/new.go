package app

import(
	ids "github.com/edemond/wiki/id"
	wikicontext "github.com/edemond/wiki/context"
	"github.com/edemond/wiki/graph"
	"github.com/edemond/wiki/log"
	"github.com/edemond/wiki/search"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

var (
	_newPageTemplate *template.Template
	_aliasForbiddenChars *regexp.Regexp
)

func init() {
	_newPageTemplate = mustParseTemplates("new.html", "new.html", "base.html")
	_aliasForbiddenChars = regexp.MustCompile("[^0-9A-Za-z- ]")
}

type NewPage struct {
	Contents string
	Title string
}

// newPage displays the editor to create a new page and handles new page submission.
func newPage(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		newPageGet(rw, r)
	} else if r.Method == "POST" {
		newOrEditPagePost(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed) 
	}
}

func newPageGet(rw http.ResponseWriter, r *http.Request) {
	ctx := wikicontext.GetHTTPContext(r)
	err := _newPageTemplate.ExecuteTemplate(rw, "base", &NewPage{
		Title: "New page",
		Contents: "",
	})
	if err != nil {
		log.Errorf(ctx, "Error executing new page template: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// newOrEditPagePost writes a new or updated page.
func newOrEditPagePost(rw http.ResponseWriter, r *http.Request) {
	ctx := wikicontext.GetHTTPContext(r)

	err := r.ParseForm()
	if err != nil {
		log.Errorf(ctx, "Error parsing form params: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var id string
	if len(r.PostForm["id"]) > 0 {
		id = r.PostForm["id"][0]
	} else {
		id, err = ids.NewID()
		if err != nil {
			log.Errorf(ctx, "Error generating a new ID: %v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	title := r.PostForm["title"][0]
	contents := graph.Markdown(r.PostForm["contents"][0])

	var alias graph.Alias
	if len(r.PostForm["alias"]) > 0 {
		alias = graph.Alias(r.PostForm["alias"][0])
	} else {
		alias = getAliasFromTitle(title)
	}

	if strings.TrimSpace(string(alias)) == "" {
		// TODO: client-side validation so this doesn't happen
		log.Errorf(ctx, "Tried to write a page with an empty alias.")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	node := &graph.Node{
		ID: id,
		Alias: alias,
		Contents: contents,
		Title: title,
	}

	ns := graph.NewNodeSource()

	// TODO: Save and load nodes by the ID, not the alias. The alias
	// should be looked up by a different service, for indirection.
	err = ns.WriteNode(ctx, graph.ID(node.Alias), node)
	if err != nil {
		log.Errorf(ctx, "Error writing page to alias '%v': %v\n", alias, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} 
	log.Printf(ctx, "Wrote page with alias '%v'.", alias)

	indexNode(ctx, node)

	http.Redirect(rw, r, fmt.Sprintf("/%v", alias), 302) // TODO: constant
}

func getAliasFromTitle(title string) graph.Alias {
	alias := _aliasForbiddenChars.ReplaceAllString(title, "") 
	alias = strings.Replace(alias, " ", "-", -1)
	alias = strings.ToLower(alias)
	return graph.Alias(alias)
}

// Index the node in the search engine.
func indexNode(ctx context.Context, node *graph.Node) {
	// TODO: This should be a hook on modification of the node,
	// not a feature of the web application.

	// TODO: This should really, really be asynchronous. We made it synchronous
	// because GAE Standard environment doesn't support background contexts.
	//go func() {
		indexer := search.NewIndexer() 

		document, err := getDocumentForNode(node)
		if err != nil {
			log.Errorf(ctx, "Error creating search document from node %v: %v", node.ID, err)
			return
		}

		log.Printf(ctx, "Indexing node with title '%v' and alias '%v'", document.Title, document.Alias)
		err = indexer.Index(ctx, document)
		if err != nil {
			log.Errorf(ctx, "Error indexing node: %v", err)
			return
		}
		log.Printf(ctx, "Indexed node with ID %v, alias '%v'.", node.ID, node.Alias)
	//}()
}

func getDocumentForNode(node *graph.Node) (*search.Document, error) {
	return &search.Document{
		Alias: string(node.Alias),
		Contents: string(node.Contents),
		ID: fmt.Sprintf("%v", node.ID),
		Title: node.Title,
	}, nil
}
