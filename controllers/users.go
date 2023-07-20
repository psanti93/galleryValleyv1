package controllers

import (
	"fmt"
	"net/http"

	"github.com/psanti93/galleryValleyv1/models"
)

type Users struct {
	Templates struct {
		SignUp View //using the view interface rather than reyling on the views.Templates package
		SignIn View
	}
	UserService    *models.UserService
	SessionService *models.SessionService
}

// Signing Up a New User
func (u Users) SignUp(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignUp.Execute(w, r, data) // passes data that get from thhe form
}

func (u Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)

	if err != nil {
		fmt.Println(err)
		// TODO: LONG term - show a warning that we're not able to sign a user in
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// Signing in Functionalities

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data) // passes data that get from thhe form
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

	session, err := u.SessionService.Create(user.ID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "User doesn't exist", http.StatusInternalServerError)
		return
	}
	// creating a cookie with golang
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true, // only want cookies to be accessible via http browser requests, don't allow cookies to work for java script
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// reading a cookie with golang
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {

	tokenCookie, err := r.Cookie("session")

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	// uses the user function from session service to look up the user based on that cookie
	user, err := u.SessionService.User(tokenCookie.Value)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "Current user: %s\n", user.Email)

}
