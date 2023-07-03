package controllers

import (
	"net/http"
)

//creating a closure that takes in a template and returns back a handler

func StaticHandler(tpl View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl View) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free 30 day trail version",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have 24/7 support staff answering emails",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@galleyvalley.com">support@galleyvalley.com</a>`,
		},
		{
			Question: "Where is your office located?",
			Answer:   "Out entire team is remote",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
