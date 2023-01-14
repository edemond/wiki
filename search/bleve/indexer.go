package bleve

import(
	"github.com/edemond/wiki/log"
	wiki "github.com/edemond/wiki/search"
	"context"
	"fmt"
)

type bleveIndexer struct {
	path string
}

func (b *bleveIndexer) Index(ctx context.Context, document *wiki.Document) error {
	index, err := openIndex(ctx, b.path)
	if err != nil {
		return err
	}
	defer closeIndex(ctx, index)

	err = index.Index(string(document.ID), document)
	if err != nil {
		return fmt.Errorf("Error indexing document: %v", err)
	}

	log.Printf(ctx, "Indexed document with ID '%v', Title '%v', Contents '%v'.", document.ID, document.Title, document.Contents)

	return nil
}
