package views

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"io/fs"
	"net/http"
)

//go:embed layouts/*.gohtml
var layouts embed.FS

var (
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, status int, files ...string) *View {
	// Prepend and append file paths with "views"
	// and ".gohtml"
	addTemplatePath(files)
	addTemplateExtension(files)

	files = append(files, layoutFiles()...)

	t, err := template.New("").ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
		Status:   status,
	}
}

type View struct {
	Template *template.Template
	Layout   string
	Status   int
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r)
}

func (v *View) Render(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(v.Status)
	w.Header().Set("Content-Type", "text/html")

	var buf bytes.Buffer

	template := v.Template.Funcs(template.FuncMap{})

	if err := template.ExecuteTemplate(&buf, v.Layout, nil); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

// Returns a slice of strings
// representing the layout files
func layoutFiles() []string {

	files, err := fs.Glob(layouts, "layouts/*.gohtml")
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		files[i] = TemplateDir + file
	}
	return files
}

// This function takes in a slice of strings
// representing files path for templates
// and it prepends the TemplateDir to each string
// in the slice.
//
// Ex.: "home" would result in "views/home/" if
// TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExtension(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
