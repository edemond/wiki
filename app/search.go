package app

import(
	wikicontext "github.com/edemond/wiki/context"
	"github.com/edemond/wiki/log"
	"github.com/edemond/wiki/search"
	"context"
	"html/template"
	"net/http"
	"net/url"
)

var _searchPageTemplate *template.Template

func init() {
	_searchPageTemplate = mustParseTemplates("search.html", "search.html", "base.html")
}

type SearchPage struct {
	Count uint64
	Query string
	Results []*SearchPageResult
}

type SearchPageResult struct {
	Title string
	Preview string
	URL *url.URL
}

func searchPage(rw http.ResponseWriter, r *http.Request) {
	ctx := wikicontext.GetHTTPContext(r)
	engine := search.NewSearch() 
	query := search.Query{
		Text: r.URL.Query()["q"][0],
	}

	log.Printf(ctx, "Searching for: '%v'", query.Text)

	results, err := engine.Search(ctx, query)
	if err != nil {
		log.Errorf(ctx, "Error searching for '%v': %v", query.Text, err)
	}

	_searchPageTemplate.ExecuteTemplate(rw, "base", getSearchPageForResults(
		ctx,
		&query,
		results,
	))
}

func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}

func getSearchPageForResults(ctx context.Context, query *search.Query, results *search.Results) *SearchPage {
	sprs := []*SearchPageResult{}
	for _,result := range results.Results {
		url, err := url.Parse(result.Alias)
		if err != nil {
			log.Errorf(ctx, "Couldn't parse URL for alias '%v'.", result.Alias)
		}

		sprs = append(sprs, &SearchPageResult{
			Title: result.Title,
			Preview: truncate(result.Preview, 250),
			URL: url,
		})
	}

	return &SearchPage{
		Count: results.Count,
		Query: query.Text,
		Results: sprs,
	}
}
