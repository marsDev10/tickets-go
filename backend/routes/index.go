package routes

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/marsDev10/helpdesk-backend/middleware"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()

	// Redirect to route with or without trailing slash
	router.StrictSlash(true)

	router.Use(loggingMiddleware)
	router.Use(middleware.CORSMiddleware(getAllowedOrigins()))

	// Handle preflight requests for any route early so CORS headers are sent.
	router.PathPrefix("/").Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

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

	// Tickets router
	ticketsRouter := protectedRouter.PathPrefix("/tickets").Subrouter()

	ticketsRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager", "agent", "member")(
			http.HandlerFunc(CreateTicketHandler),
		),
	).Methods("POST")

	ticketsRouter.Handle("/{ticket_id}",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(DeleteTicketHandler),
		),
	).Methods("DELETE")

	ticketsRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager", "agent", "member", "viewer")(
			http.HandlerFunc(ListTicketsHandler),
		),
	).Methods("GET")

	ticketsRouter.Handle("/{ticket_id}",
		middleware.RoleMiddleware("admin", "manager", "agent", "member", "viewer")(
			http.HandlerFunc(GetTicketHandler),
		),
	).Methods("GET")

	ticketsRouter.Handle("/{ticket_id}",
		middleware.RoleMiddleware("admin", "manager", "agent")(
			http.HandlerFunc(UpdateTicketHandler),
		),
	).Methods("PATCH")

	ticketsRouter.Handle("/{ticket_id}/assign",
		middleware.RoleMiddleware("admin", "manager", "supervisor")(
			http.HandlerFunc(AssignTicketHandler),
		),
	).Methods("POST")

	teamRouter := protectedRouter.PathPrefix("/teams").Subrouter()

	teamRouter.Handle("/",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(CreateTeamHandler),
		),
	).Methods("POST")

	teamRouter.Handle("/{team_id}",
		middleware.RoleMiddleware("admin", "manager")(
			http.HandlerFunc(UpdateTeamHandler),
		),
	).Methods("PUT")

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

func getAllowedOrigins() []string {
	raw := strings.TrimSpace(os.Getenv("CORS_ALLOWED_ORIGINS"))
	if raw == "" {
		return []string{"*"}
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))

	for _, part := range parts {
		if origin := strings.TrimSpace(part); origin != "" {
			origins = append(origins, origin)
		}
	}

	if len(origins) == 0 {
		return []string{"*"}
	}

	return origins
}
