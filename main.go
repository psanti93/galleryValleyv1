package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/psanti93/galleryValleyv1/controllers"
	"github.com/psanti93/galleryValleyv1/models"
	"github.com/psanti93/galleryValleyv1/templates"
	"github.com/psanti93/galleryValleyv1/views"
)

func main() {

	r := chi.NewRouter()

	// parsing the template prior to executing it
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	//Contact
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	// FAQ
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	// Configuring the users controller to include the user service
	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userService := models.UserService{
		DB: db}

	sessionService := models.SessionService{
		DB: db,
	}

	usersC := controllers.Users{UserService: &userService, SessionService: &sessionService}

	// Sign Up
	usersC.Templates.SignUp = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.SignUp)
	r.Post("/users", usersC.CreateUser)

	// Sign In
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)

	// getting the cookie of the user
	r.Get("/users/me", usersC.CurrentUser)

	//Add sign outrout
	r.Post("/signout", usersC.ProcessSignOut)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 3000....")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"               // needs a 32 character auth key
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false)) // csrf.Secure() by default it's true, it requires an https secure connection, false for now cause local we don't have https connection. set to true in prod
	http.ListenAndServe(":3000", csrfMw(r))

}
