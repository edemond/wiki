package app

import(
	"html/template"
	"os"
	"path/filepath"
)

func getDirectory(envVariable string) string {
	dir, ok := os.LookupEnv(envVariable) 
	if !ok {
		return "."
	}
	return dir
}

// Get the directory where templates are stored.
func getTemplateDirectory() string {
	return getDirectory("WIKI_TEMPLATE_DIR")
}

func getStaticDirectory() string {
	return getDirectory("WIKI_STATIC_DIR")
}

func mustParseTemplates(name string, files ...string) *template.Template {
	dir := getTemplateDirectory()

	paths := []string{}
	for _,file := range files {
		paths = append(paths, filepath.Join(dir, file))
	}

	return template.Must(template.New(name).ParseFiles(paths...))
}