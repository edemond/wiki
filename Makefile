appengine:
	go build github.com/edemond/wiki/cmd/wiki-appengine

plain:
	go build github.com/edemond/wiki/cmd/wiki-plain

all: appengine plain
