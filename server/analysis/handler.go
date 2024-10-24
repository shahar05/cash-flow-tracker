package analysis

import (
	"database/sql"
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
	r.HandleFunc("/analysis", GetAnalysisHandler).Methods("GET")
	r.HandleFunc("/category-analysis", GetCategoryAnalysisHandler).Methods("GET")
	r.HandleFunc("/category-graph", GetCategoryGraphHandler).Methods("GET")
	r.HandleFunc("/categories-graph", GetCategoriesGraphHandler).Methods("GET")
}

func GetCategoriesGraphHandler(w http.ResponseWriter, r *http.Request) {
	response, err := GetMonthlyTransactions()
	if err != nil {
		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSONOk(w, response)
}

func GetCategoryGraphHandler(w http.ResponseWriter, r *http.Request) {
	response, err := GetMonthlyTransactionSumsByCategory(r.URL.Query().Get("name"))
	if err != nil {
		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSONOk(w, response)
}

//

func GetCategoryAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	response, err := GetCategoryAnalysis(r.URL.Query().Get("name"))
	if err != nil {
		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSONOk(w, response)
}

func GetAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	// var dates DateParams
	// if err := json.NewDecoder(r.Body).Decode(&dates); err != nil {
	// 	log.Printf("GetAnalysisHandler: Error decoding transaction: %v", err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	response, err := GetCategorySums()
	if err != nil {
		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSONOk(w, response)
}
