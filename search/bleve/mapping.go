package bleve 

import(
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/lang/en"
	"github.com/blevesearch/bleve/mapping"
)

func getPageDocumentMapping() *mapping.DocumentMapping {
	numericField := bleve.NewNumericFieldMapping()

	englishField := bleve.NewTextFieldMapping()
	englishField.Analyzer = en.AnalyzerName

	textField := bleve.NewTextFieldMapping()

	pageMapping := bleve.NewDocumentMapping()
	pageMapping.AddFieldMappingsAt("Id", numericField)
	pageMapping.AddFieldMappingsAt("Title", englishField)
	pageMapping.AddFieldMappingsAt("Alias", textField)
	pageMapping.AddFieldMappingsAt("Contents", englishField) // TODO: Custom analyzer for Markdown?

	return pageMapping
}

/*
func getPersonDocumentMapping() *mapping.DocumentMapping {
	panic("todo")
}
*/

func getIndexMapping() (mapping.IndexMapping, error) {
	mapping := bleve.NewIndexMapping()
	mapping.AddDocumentMapping("page", getPageDocumentMapping())
	/*
	//mapping.AddDocumentMapping("person", getPersonDocumentMapping())
	*/
	return mapping, nil
}