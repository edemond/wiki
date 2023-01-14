package app

import(
	"github.com/edemond/wiki/graph"
	"github.com/golang-commonmark/markdown"
)

func getHTMLFromMarkdown(m graph.Markdown) (string, error) {
	md := markdown.New(
		markdown.Linkify(true),
	)
	return md.RenderToString([]byte(m)), nil
}
