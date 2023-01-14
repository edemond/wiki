package graph

import(
	"context"
)

type Alias string // "Friendly name" for a node.
type ID string // ID for an object in the graph.
type Markdown string

type Node struct {
	ID string 
	Alias Alias // TODO: Remove this from Node
	Contents Markdown
	Title string
}

// Stores data for nodes in the graph.
type NodeSource interface {
	GetNode(ctx context.Context, id ID) (*Node, error)
	WriteNode(ctx context.Context, id ID, n *Node) error
}

// Stores the graph.
type Graph interface {
	// TODO: Hmm, what about query languages here?
	AddEdge(from ID, edge string, to ID) error
	QueryEdge(from ID, edge string) ([]ID, error)
	Close() error
}

func NewNodeSource() NodeSource {
	driver := getNodeDriver()
	return driver.NewNodeSource()
}

func NewGraph() Graph {
	driver := getGraphDriver()
	return driver.NewGraph()
}
