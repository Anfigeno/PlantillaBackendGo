package controladores

import "net/http"

type Controlador struct {
	Get    func(w http.ResponseWriter, r *http.Request)
	Post   func(w http.ResponseWriter, r *http.Request)
	Put    func(w http.ResponseWriter, r *http.Request)
	Delete func(w http.ResponseWriter, r *http.Request)
}
