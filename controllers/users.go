package controllers

import (
	"fmt"
	"net/http"

	"github.com/psanti93/galleryValleyv1/context"
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

	setCookie(w, CookieSession, session.Token)
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

	// setting cookie
	setCookie(w, CookieSession, session.Token)

	http.Redirect(w, r, "/users/me", http.StatusFound)
}

// reading a cookie with golang
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {

	// Checks if a user currently exists with the request
	user := context.User(r.Context())

	fmt.Fprintf(w, "Current user: %s\n", user.Email)

}

func (u Users) ProcessSignOut(w http.ResponseWriter, r *http.Request) {
	token, err := readCookie(r, CookieSession)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	err = u.SessionService.Delete(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	// TOdo: Delete the user's coookie
	deleteCookie(w, CookieSession)

	http.Redirect(w, r, "/signin", http.StatusFound)
}

//User Middleware

type UserMiddleware struct {
	SessionService *models.SessionService
}

func (umw UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := readCookie(r, CookieSession)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user, err := umw.SessionService.User(token)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}

func (umw UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)

	})
}
