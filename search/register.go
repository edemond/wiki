package search

import(
)

var _searchDriver SearchDriver
var _indexerDriver IndexerDriver

type SearchDriver interface {
	NewSearch() Search
}

type IndexerDriver interface {
	NewIndexer() Indexer
}

func getSearchDriver() SearchDriver {
	if _searchDriver == nil {
		panic("No search driver registered!")
	}
	return _searchDriver
}

func getIndexerDriver() IndexerDriver {
	if _indexerDriver == nil {
		panic("No indexer driver registered!")
	}
	return _indexerDriver
}

func RegisterSearchDriver(driver SearchDriver) {
	if _searchDriver != nil {
		panic("Search driver registered twice!")
	} else if driver == nil {
		panic("Tried to register a nil search driver!")
	}
	_searchDriver = driver
}

func RegisterIndexerDriver(driver IndexerDriver) {
	if _indexerDriver != nil {
		panic("Indexer driver registered twice!")
	} else if driver == nil {
		panic("Tried to register a nil indexer driver!")
	}
	_indexerDriver = driver
}