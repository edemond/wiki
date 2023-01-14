package bleve

import(
	"github.com/edemond/wiki/log"
	wiki "github.com/edemond/wiki/search"
	"context"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"github.com/blevesearch/bleve/analysis/lang/en"
)

const _INDEX_NAME string = "index.bleve" // TODO: Config this

func init() {
	wiki.RegisterSearchDriver(&bleveSearchDriver{})
	wiki.RegisterIndexerDriver(&bleveIndexerDriver{})
}

type bleveSearchDriver struct {}

func (b *bleveSearchDriver) NewSearch() wiki.Search {
	return &bleveSearch{
		path: _INDEX_NAME,
	}
}

type bleveIndexerDriver struct {}

func (b *bleveIndexerDriver) NewIndexer() wiki.Indexer {
	return &bleveIndexer{
		path: _INDEX_NAME,
	}
}

type bleveSearch struct {
	path string
}

func createIndex(ctx context.Context, path string) (bleve.Index, error) {
	log.Printf(ctx, "Creating index at '%v'...", path)
	mapping, err := getIndexMapping()
	if err != nil {
		return nil, fmt.Errorf("Couldn't create mapping: %v", err)
	}

	index, err := bleve.New(path, mapping)
	if err != nil {
		return nil, fmt.Errorf("Error creating index at '%v': %v", path, err)
	}
	return index, err
}

func openIndex(ctx context.Context, path string) (bleve.Index, error) {
	index, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf(ctx, "Index does not exist at '%v'. Creating it.", path)
		return createIndex(ctx, path)
	} else if err != nil {
		return nil, fmt.Errorf("Error opening index at '%v': %v", path, err)
	}
	return index, nil
}

func closeIndex(ctx context.Context, index bleve.Index) {
	err := index.Close()
	if err != nil {
		log.Errorf(ctx, "Error closing index: '%v'", err)
	}
}

func (b *bleveSearch) Search(ctx context.Context, query wiki.Query) (*wiki.Results, error) {
	index, err := openIndex(ctx, b.path)
	if err != nil {
		return nil, err
	}
	defer closeIndex(ctx, index)

	q := bleve.NewMatchQuery(query.Text)
	q.Analyzer = en.AnalyzerName // TODO: Is this the right way to specify an analyzer?
	request := bleve.NewSearchRequest(q)
	request.Fields = []string{"Title", "Contents", "Alias"}

	bleveResults, err := index.Search(request)
	if err != nil {
		return nil, fmt.Errorf("Error running search query '%v': %v", query.Text, err)
	}

	// TODO: Lots of other good stuff in SearchResult.
	log.Printf(
		ctx,
		"Search query for '%v' returned %v results and took %v.", 
		query.Text, 
		bleveResults.Total, 
		bleveResults.Took,
	)

	results := []*wiki.Result{}
	for _,hit := range bleveResults.Hits {
		result, err := getResultForHit(hit)
		if err != nil {
			log.Errorf(ctx, "%v", err)
			continue
		}
		results = append(results, result)
	}

	return &wiki.Results{
		Count: bleveResults.Total,
		Results: results,
	}, nil
}

// Get a wiki search result for a Bleve search engine hit.
func getResultForHit(hit *search.DocumentMatch) (*wiki.Result, error) {
	// TODO: Why are these capitalized?! It should be "title" and "contents"
	title, ok := hit.Fields["Title"].(string)
	if !ok {
		return nil, fmt.Errorf("Can't convert title '%v' to string.", hit.Fields["Title"])
	}

	preview, ok := hit.Fields["Contents"].(string)
	if !ok {
		return nil, fmt.Errorf("Can't convert contents '%v' to string.", hit.Fields["Contents"])
	}

	alias, ok := hit.Fields["Alias"].(string)
	if !ok {
		return nil, fmt.Errorf("Can't convert alias '%v' to string.", hit.Fields["Alias"])
	}

	return &wiki.Result{
		ID: hit.ID,
		Title: title,
		Preview: preview, // TODO: how do we get the matching parts out? Fragments?
		Alias: alias,
	}, nil
}
