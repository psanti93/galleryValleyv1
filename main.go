package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/psanti93/galleryValleyv1/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

//Create Helper Function to execute the template

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)

	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
