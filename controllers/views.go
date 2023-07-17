package controllers

import "net/http"

type View interface {
	Execute(w http.ResponseWriter, r *http.Request, data interface{})
}
