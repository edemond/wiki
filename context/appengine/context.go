// Package context provides a source of context.Context
// that uses appengine.NewContext.
package appengine

import(
	wiki "github.com/edemond/wiki/context"
	"context"
	"google.golang.org/appengine"
	"net/http"
)

func init() {
	wiki.Register(&source{})
}

type source struct {}

func (s *source) GetHTTPContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}
