package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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
			"csrfField": func() template.HTML {
				return `<!-- TODO implement CSRF Field -->` // 1. Parses a filler function that will later be filled in Execute Line 46
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
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Errorf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
