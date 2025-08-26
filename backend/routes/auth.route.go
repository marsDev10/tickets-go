package routes

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar el login
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("LoginHandler no implementado"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para manejar el registro
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("RegisterHandler no implementado"))
}
