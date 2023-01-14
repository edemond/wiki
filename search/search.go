/*
	Package search implements a search engine interface for the wiki.

	Multiple drivers are supported, including Bleve and Google App
	Engine search.

	Document types to support:
	- Page
	- Person
	- more to come
*/
package search

import(
	"context"
)

// Search represents a connection to the search engine for querying.
type Search interface {
	Search(ctx context.Context, query Query) (*Results, error)
}

// Indexer
type Indexer interface {
	Index(ctx context.Context, document *Document) error 
}

// Query represents a search engine query.
type Query struct {
	Text string
}

// Results holds the results of a search query.
type Results struct {
	Count uint64
	Results []*Result
}

// Result is one search query result.
type Result struct {
	ID string
	Alias string // TODO: Is it a mistake to have the search engine know about alias? Is that too specific?
	Title string
	Preview string
	// TODO: Highlighted matches
}

// TODO: Why would we not just reuse Document inside Result?
// Document is a search index document.
type Document struct {
	ID string
	Title string
	Contents string
	Alias string
}

// Implement mapping.Clasifier
// TODO: We eventually want more than one search document type.
func (d *Document) Type() string {
	return "page" // Needs to agree with name of doctype in mapping
}

func NewSearch() Search {
	driver := getSearchDriver()
	return driver.NewSearch()
}

func NewIndexer() Indexer {
	driver := getIndexerDriver()
	return driver.NewIndexer()
}