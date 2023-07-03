package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		View View //using the view interface rather than reyling on the views.Templates package
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// we need a view to render
	u.Templates.View.Execute(w, nil)
}
