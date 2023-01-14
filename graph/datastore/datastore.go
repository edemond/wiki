// Package datastore implements node storage on Google App Engine datstore.
package datastore

import(
	"github.com/edemond/wiki/graph"
	"context"
	"google.golang.org/appengine/datastore"
)

const _KIND string = "node"

type source struct {}

type driver struct {}

func (d *driver) NewNodeSource() graph.NodeSource {
	return &source{}
}

func init() {
	graph.RegisterNodeDriver(&driver{})	
}

func (s *source) GetNode(ctx context.Context, id graph.ID) (*graph.Node, error) {
	key := datastore.NewKey(ctx, _KIND, string(id), 0, nil)

	var node graph.Node
	err := datastore.Get(ctx, key, &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func (s *source) WriteNode(ctx context.Context, id graph.ID, node *graph.Node) error {
	key := datastore.NewKey(ctx, _KIND, string(id), 0, nil)

	var err error
	key, err = datastore.Put(ctx, key, node)
	if err != nil {
		return err
	}

	return nil
}
