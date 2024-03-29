package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jgsheppa/jgsheppa.github.io/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	view := controllers.NewUser()

	// Index handler
	r.Get("/", view.ProfileView.Render)
	r.NotFound(view.NotFoundView.Render)

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filesDir := http.Dir(filepath.Join(workDir, "images"))
	FileServer(r, "/images", filesDir)

	http.ListenAndServe(":3000", r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
