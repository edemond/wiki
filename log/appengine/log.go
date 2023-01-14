// Package appengine implements logging on Google App Engine.
package appengine

import(
	"google.golang.org/appengine/log"
	wiki "github.com/edemond/wiki/log"
	"context"
)

func init() {
	wiki.Register(&logger{})
}

type logger struct {}

func (l *logger) Printf(ctx context.Context, fmt string, args ...interface{}) {
	log.Infof(ctx, fmt, args...)
}

func (l *logger) Errorf(ctx context.Context, fmt string, args ...interface{}) {
	log.Errorf(ctx, fmt, args...)
}
