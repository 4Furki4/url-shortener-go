package main

import (
	"log"
	"net/http"
	"url-shortener/database"
	"url-shortener/routes"
)

func main() {
	database.InitDatabase()   // Initialize SQLite
	defer database.DB.Close() // Ensure database is closed on exit
	router := routes.SetupRoutes()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
