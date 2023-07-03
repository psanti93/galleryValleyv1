package controllers

import (
	"net/http"

	"github.com/psanti93/galleryValleyv1/views"
)

type Users struct {
	Templates struct {
		View views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// we need a view to render
	u.Templates.View.Execute(w, nil)
}
