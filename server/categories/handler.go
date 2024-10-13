package categories

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shahar05/cash-flow-viewer/utils"
)

var DB *sql.DB

// SetDB sets the database connection
func SetDB(db *sql.DB) {
	DB = db
}

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/categories", GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories", AddCategoriesHandler).Methods("POST")
}

// GetCategoriesHandler handles GET requests for contacts with pagination
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCategoriesHandler: Handling POST /Categories request")

	cats, err := GetCategories()
	if err != nil {
		log.Printf("GetCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("GetCategoriesHandler: sending categories")
	utils.WriteJSONOk(w, cats)
}

// AddCategoriesHandler handles POST requests to add a new transaction
func AddCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AddCategoriesHandler: Handling POST /Categories request")
	var categoryArr CategoryArray
	if err := json.NewDecoder(r.Body).Decode(&categoryArr); err != nil {
		log.Printf("AddCategoriesHandler: Error decoding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	successfullyInserted, err := AddCategories(categoryArr.Categories)
	if err != nil {
		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("AddCategoriesHandler: Added %d new categories", successfullyInserted)
	utils.WriteJSONOk(w, map[string]int64{"successfullyInserted": successfullyInserted})
}
