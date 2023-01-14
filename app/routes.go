package app

import(
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

// GetRouter returns an http.Handler to handle all Wiki HTTP requests.
func GetRouter(serveStaticFiles bool) http.Handler {
	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/search", searchPage)
	r.HandleFunc("/new", newPage)
	if serveStaticFiles {
		static := http.Dir(getStaticDirectory())
		r.PathPrefix("/css/").Handler(http.FileServer(static))
		r.PathPrefix("/js/").Handler(http.FileServer(static))
	}
	r.HandleFunc("/edit/{alias}", editPage)
	r.HandleFunc("/edit/", editPage)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.HandleFunc("/{alias}", readPage)

	// Middleware
	n := negroni.New() 
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)

	return n
}
