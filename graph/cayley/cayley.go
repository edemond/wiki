// Package cayley implements the Wiki graph API
// using the graph database Cayley.
package cayley

import (
	wiki "github.com/edemond/wiki/graph"
	"fmt"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/kv/bolt"
	"github.com/cayleygraph/cayley/quad"
)

func init() {
	wiki.RegisterGraphDriver(&driver{})
}

type driver struct {}

func (c *driver) NewGraph() wiki.Graph {
	return &cayleyGraph{
		path: "graph.cayley",
	}
}

type cayleyGraph struct {
	path string
	handle *cayley.Handle
}

func openCayleyGraph(filename string) (wiki.Graph, error) {
	// TODO: Does this overwrite an old graph, or open it?!
	// TODO: What is the third argument to InitQuadStore?
	err := graph.InitQuadStore("bolt", filename, nil)

	handle, err := cayley.NewGraph("bolt", filename, nil)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open Cayley graph at '%v': %v", filename, err)
	}

	return &cayleyGraph{
		handle: handle,
	}, nil
}

func (g *cayleyGraph) Close() error {
	return g.handle.Close()
}

func (g *cayleyGraph) AddEdge(from wiki.ID, edge string, to wiki.ID) error {
	return g.handle.AddQuad(quad.Make(from, edge, to, nil))
}

func (g *cayleyGraph) QueryEdge(from wiki.ID, edge string) ([]wiki.ID, error) {
	panic("TODO")
}
