package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/psanti93/galleryValleyv1/controllers"
	"github.com/psanti93/galleryValleyv1/views"
)

func main() {

	r := chi.NewRouter()

	// parsing the template prior to executing it
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}

	//comparison between static handler and using a regular handler func
	r.Get("/", controllers.StaticHandler(tpl))

	// Contact
	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	// FAQ
	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
