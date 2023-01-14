package log

import(
	"context"
	"log"
)

// Logger defines a logging interface for the Wiki.
type Logger interface {
	Printf(ctx context.Context, fmt string, args ...interface{})
	Errorf(ctx context.Context, fmt string, args ...interface{})
}

var _defaultLogger Logger = &defaultLogger{}
var _logger Logger = _defaultLogger

func Register(logger Logger) {
	if _logger != _defaultLogger {
		panic("Logger already registered!")
	} else if logger == nil {
		panic("Tried to register a nil logger!")
	}
	_logger = logger
}

func Printf(ctx context.Context, fmt string, args ...interface{}) {
	_logger.Printf(ctx, fmt, args...)
}

func Errorf(ctx context.Context, fmt string, args ...interface{}) {
	_logger.Errorf(ctx, fmt, args...)
}

// Type default implements a logger that just uses the plain Go "log" package.
type defaultLogger struct {}

func (d *defaultLogger) Printf(_ context.Context, fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

func (d *defaultLogger) Errorf(_ context.Context, fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}

