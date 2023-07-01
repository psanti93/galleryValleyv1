package controllers

import (
	"net/http"

	"github.com/psanti93/galleryValleyv1/views"
)

//creating a closure that takes in a template and returns back a handler

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
