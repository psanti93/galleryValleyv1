package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
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

	usersC := controllers.Users{UserService: &userService}

	usersC.Templates.View = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	//testing out parsing sign up form
	r.Post("/users", usersC.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on port 3000....")

	http.ListenAndServe(":3000", r)
}
