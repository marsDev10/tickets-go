package main

import (
	"log"
	"net/http"
	"os"

	"github.com/marsDev10/helpdesk-backend/db"
	"github.com/marsDev10/helpdesk-backend/models"
	"github.com/marsDev10/helpdesk-backend/routes"
)

func main() {
	log.Printf("🚀 Iniciando el servidor...")

	db.DBConnection()
	db.DB.AutoMigrate(
		&models.Organization{},
		&models.User{},
		&models.Ticket{},
		&models.Team{},
		&models.TeamMember{},
	)

	router := routes.InitRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en el puerto %s 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
