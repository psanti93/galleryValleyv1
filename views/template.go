package views

import (
	"fmt"
	"net/http"
	"text/template"
)

type Template struct {
	view *template.Template
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %w", err)
	}
	return Template{view: tpl}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.view.Execute(w, nil)
	if err != nil {
		fmt.Errorf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
