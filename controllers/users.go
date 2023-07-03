package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		View View //using the view interface rather than reyling on the views.Templates package
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	u.Templates.View.Execute(w, data) // passes data that get from thhe form
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
