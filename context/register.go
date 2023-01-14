// Package context abstracts over different sources of
// context.Context, which have different implementations
// for e.g. App Engine vs. vanilla net/http.
//
// If no sources are registered, a default source will 
// be used.
package context

import(
	"context"
	"net/http"
)

var _defaultSource *defaultSource = &defaultSource{}
var _source contextSource = _defaultSource

type contextSource interface {
	GetHTTPContext(r *http.Request) context.Context
}

func Register(source contextSource) {
	if _source != _defaultSource {
		panic("Already registered a contextSource")
	} else if source == nil {
		panic("Tried to register a nil contextSource")
	}
	_source = source
}

func GetHTTPContext(r *http.Request) context.Context {
	return _source.GetHTTPContext(r)
}

type defaultSource struct {}

// Default implementation of GetHTTPContext, which just returns
// the regular context out of the HTTP request.
func (s *defaultSource) GetHTTPContext(r *http.Request) context.Context {
	return r.Context()
}
