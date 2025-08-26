package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	authRouter := router.PathPrefix("/api/auth").Subrouter()
	// Aquí puedes agregar las rutas de autenticación
	authRouter.HandleFunc("/login", LoginHandler).Methods("POST")
	authRouter.HandleFunc("/register", RegisterHandler).Methods("POST")

	return router

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completado en %v", time.Since(start))
	})
}
