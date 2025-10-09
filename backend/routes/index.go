package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/marsDev10/helpdesk-backend/middleware"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	authRouter := router.PathPrefix("/api/auth").Subrouter()
	// Aquí puedes agregar las rutas de autenticación
	authRouter.HandleFunc("/login", LoginHandler).Methods("POST")
	authRouter.HandleFunc("/register", RegisterHandler).Methods("POST")

	protectedRouter := router.PathPrefix("/api").Subrouter()

	protectedRouter.Use(middleware.JWTMiddleware)

	organitationRouter := protectedRouter.PathPrefix("/organitation").Subrouter()

	organitationRouter.HandleFunc("/register", CreateOrganizationHandler).Methods("POST")

	usersRouter := protectedRouter.PathPrefix("/users").Subrouter()

	usersRouter.Handle("/organization",
		middleware.RoleMiddleware("admin")(
			http.HandlerFunc(GetOrganizationUsersHandler),
		),
	).Methods("GET")

	usersRouter.Handle("/organization/{idUser}",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(GetOrganizationUserHandler),
		),
	).Methods("GET")

	usersRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(CreateUserHandler),
		),
	).Methods("POST")

	usersRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(UpdateUserHandler),
		),
	).Methods("PUT")

	usersRouter.Handle("/status",
		middleware.RoleMiddleware("admin")(
			http.HandlerFunc(ToggleUserStatusHandler),
		),
	).Methods("POST")

	teamRouter := protectedRouter.PathPrefix("/teams").Subrouter()

	teamRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(CreateTeamHandler),
		),
	).Methods("POST")

    teamRouter.Handle("/{team_id}/members",
        middleware.RoleMiddleware("admin", "manager")(
            http.HandlerFunc(AddMemberToTeamHandler),
        ),
    ).Methods("POST")

	teamRouter.Handle("/all",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(GetTeamsByOrganizationHandler),
		),
	).Methods("GET")

    teamRouter.Handle("/{team_id}/members/{user_id}",
        middleware.RoleMiddleware("admin", "manager")(
            http.HandlerFunc(RemoveMemberFromTeamHandler),
        ),
    ).Methods("DELETE")

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
