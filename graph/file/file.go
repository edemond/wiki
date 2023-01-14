package file

import(
	"github.com/edemond/wiki/graph"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type source struct {}

type driver struct {}

func (f *driver) NewNodeSource() graph.NodeSource {
	return &source{}	
}

func init() {
	graph.RegisterNodeDriver(&driver{})	
}

func getFilenameFromID(id graph.ID) string {
	// TODO: make this no longer the dumbest possible thing
	return fmt.Sprintf("nodes/%v.json", id)
}

func (s *source) GetNode(_ context.Context, id graph.ID) (*graph.Node, error) {
	filename := getFilenameFromID(id)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(
			"Couldn't get node with filename '%v' for ID '%v': %v", 
			filename, id, err,
		)
	}

	var node graph.Node 
	err = json.Unmarshal(bytes, &node)
	if err != nil {
		return nil, fmt.Errorf(
			"Couldn't unmarshal node with ID '%v' from JSON: %v", 
			id, 
			err,
		)
	}

	return &node, nil
}

func (s *source) WriteNode(_ context.Context, id graph.ID, n *graph.Node) error {
	bytes, err := json.Marshal(n)
	if err != nil {
		return fmt.Errorf(
			"Couldn't marshal node with ID '%v' to JSON: %v", 
			id,
			err,
		)
	}

	filename := getFilenameFromID(id) 
	return ioutil.WriteFile(filename, bytes, 0) // TODO: FileMode bits?
}
