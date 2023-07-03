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
	// we need a view to render
	u.Templates.View.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	/** one way to retrieve from input fields
		fmt.Fprint(w, "Email: ", r.PostForm.Get("email"))
	    fmt.Fprint(w, "Password: ", r.PostForm.Get("password"))

	**/

	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
