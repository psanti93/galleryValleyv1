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

	// creating a cookie with golang
	cookie := http.Cookie{
		Name:     "CookiePracticeEmail",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true, // only want cookies to be accessible via http browser requests, don't allow cookies to work for java script
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User Authenticated: %+v", user)
}

// reading a cookie with golang
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("CookiePracticeEmail")

	if err != nil {
		fmt.Fprint(w, "The email cookie could not be read")
		return
	}

	fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)

}
