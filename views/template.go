package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/psanti93/galleryValleyv1/context"
	"github.com/psanti93/galleryValleyv1/models"
)

type Template struct {
	view *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0]) // the first gohtml page that gets passed in
	// adding the function called csrfField to the instantiated template
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented") // 1. Parses a filler function that will later be filled in Execute Line 61
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("current user not implemented")
			},
		},
	)

	// running the parsing function against the tpl type you instantiate on line 23
	// this allows the csrfField function to be populated
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing fs template: %w", err)
	}

	return Template{view: tpl}, nil
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := t.view.Clone() // Clone() prevents race condtion --> avoids the use case of when you have multiple users, and getting the same csrfTemplate
	if err != nil {
		log.Printf("cloning template:%v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}
	//2. takes the filler function we Parsed and execute the logic for passing in a csrf token
	tpl = tpl.Funcs(
		template.FuncMap{
			// comment out to see error message example
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r) // creates the hidden token and key
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
		},
	)
	// create a buffer to execute a template in
	/** this solves the following use case from populating in the console whenever there is an error:
		http: superfluous response.WriteHeader call from github.com/psanti93/galleryValleyv1/views.Template.Execute (template.go:65)
	**/
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		fmt.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf) // copy everything from the buffer into the writer response will return a 500 status code as expected when ther eis an error
}
