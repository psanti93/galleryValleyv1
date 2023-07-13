package controllers

import (
	"fmt"
	"net/http"

	"github.com/psanti93/galleryValleyv1/models"
)

type Users struct {
	Templates struct {
		New    View //using the view interface rather than reyling on the views.Templates package
		SignIn View
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data) // passes data that get from thhe form
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// +v adds fields for users
	fmt.Fprintf(w, "User Created:%+v ", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data) // passes data that get from thhe form
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "User doesn't exist", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User Authenticated: %+v", user)
}
