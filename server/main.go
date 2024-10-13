package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/shahar05/cash-flow-viewer/categories"
	"github.com/shahar05/cash-flow-viewer/database"
	"github.com/shahar05/cash-flow-viewer/transactions"

	"github.com/gorilla/mux"
)

func main() {

	// Init DB
	db := database.Init()

	// Init Router
	r := mux.NewRouter()

	// Set up CORS with allowed origins
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500", "http://localhost:5500"}, // Allow only this origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Specify allowed methods
		AllowedHeaders:   []string{"Content-Type"},                            // Specify allowed headers
	})

	// Register the HealthCheckHandler
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	// Register the transactions party
	transactions.RegisterRoutes(r)
	transactions.SetDB(db)

	categories.RegisterRoutes(r)
	categories.SetDB(db)

	handler := c.Handler(r)
	// Start server
	portServer := "8080"
	log.Printf("Server is running on port %s", portServer)
	log.Fatal(http.ListenAndServe(":"+portServer, handler))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Check ok"))
}
