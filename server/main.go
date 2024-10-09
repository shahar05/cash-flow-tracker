package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
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

	// Register the HealthCheckHandler
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	// Register the transactions party
	transactions.RegisterRoutes(r)
	transactions.SetDB(db)

	categories.RegisterRoutes(r)
	categories.SetDB(db)

	// Start server
	portServer := "8080"
	log.Printf("Server is running on port %s", portServer)
	log.Fatal(http.ListenAndServe(":"+portServer, r))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Check ok"))
}
