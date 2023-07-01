package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/psanti93/galleryValleyv1/controllers"
	"github.com/psanti93/galleryValleyv1/templates"
	"github.com/psanti93/galleryValleyv1/views"
)

func main() {

	r := chi.NewRouter()

	// parsing the template prior to executing it
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	//Contact
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	// FAQ
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
