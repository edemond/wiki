package graph

import(
)

var (
	_nodeDriver NodeDriver
	_graphDriver GraphDriver 
)

// NodeDriver represents an implementation of the Wiki node API.
type NodeDriver interface {
	NewNodeSource() NodeSource
}

// GraphDriver represents an implementation of the Wiki graph API.
type GraphDriver interface {
	NewGraph() Graph
}

func getNodeDriver() NodeDriver {
	if _nodeDriver == nil {
		panic("No node driver registered!")
	}
	return _nodeDriver
}

func getGraphDriver() GraphDriver {
	if _graphDriver == nil {
		panic("No graph driver registered!")
	}
	return _graphDriver
}

// RegisterNodeDriver lets a package register an implementation of NodeSource.
func RegisterNodeDriver(driver NodeDriver) {
	if _nodeDriver != nil {
		panic("Node driver registered twice!")
	} else if driver == nil {
		panic("Tried to register a nil node driver")
	}
	_nodeDriver = driver
}

// RegisterGraphDriver lets a package register an implementation of GraphSource.
func RegisterGraphDriver(driver GraphDriver) {
	if _graphDriver != nil {
		panic("Graph driver registered twice")
	} else if driver == nil {
		panic("Tried to register a nil graph driver")
	}
	_graphDriver = driver
}