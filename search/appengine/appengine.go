package appengine

import(
	wiki "github.com/edemond/wiki/search"
	"context"
	"fmt"
	aesearch "google.golang.org/appengine/search"
)

type searchDriver struct {}
type indexerDriver struct {}

type search struct {
	indexName string
}
type indexer struct {
	indexName string
}

func init() {
	wiki.RegisterSearchDriver(&searchDriver{})
	wiki.RegisterIndexerDriver(&indexerDriver{})
}

func (b *searchDriver) NewSearch() wiki.Search {
	return &search{
		indexName: "nodes",
	}
}

func (b *indexerDriver) NewIndexer() wiki.Indexer {
	return &indexer{
		indexName: "nodes",
	}
}

func (s *search) Search(ctx context.Context, query wiki.Query) (*wiki.Results, error) {
	index, err := aesearch.Open(s.indexName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open search index '%v': %v", s.indexName, err)
	}

	var results wiki.Results

	// TODO: wtf is the extra semicolon
	for iter := index.Search(ctx, query.Text, nil); ; {
		var document wiki.Document
		_, err := iter.Next(&document) // TODO: Need the id for anything?
		if err == aesearch.Done {
			break
		} else if err != nil {
			// TODO: Log more about the query here?
			return nil, fmt.Errorf("Error iterating search results: %v", err)
		}

		results.Results = append(results.Results, getResultFromDocument(&document)) 
	}

	results.Count = uint64(len(results.Results))

	return &results, nil
}

func (i *indexer) Index(ctx context.Context, document *wiki.Document) error {
	index, err := aesearch.Open(i.indexName)
	if err != nil {
		return fmt.Errorf("Couldn't open search index '%v': %v", i.indexName, err)
	}

	_, err = index.Put(ctx, document.ID, document) // TODO: Have we already generated an ID?
	if err != nil {
		return fmt.Errorf("Couldn't put document: %v", err)
	}

	return nil
}

func getResultFromDocument(doc *wiki.Document) *wiki.Result {
	return &wiki.Result{
		ID: doc.ID,
		Alias: doc.Alias,
		Preview: doc.Contents, // TODO: Shorten
		Title: doc.Title,
	}
}
