package analysis

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shahar05/cash-flow-viewer/core"
)

var DB *sql.DB

// SetDB sets the database connection
func SetDB(db *sql.DB) {
	DB = db
}

// RegisterRoutes sets up the HTTP routes for contacts
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/analysis-pie-date", GetPieByDateHandler).Methods("GET")
	r.HandleFunc("/analysis-metric", GetMetricHandler).Methods("GET")
	// r.HandleFunc("/analysis-pie-monthly", GetPieMonthlyHandler).Methods("POST")
	// r.HandleFunc("/analysis-graph-monthly", GetGraphMonthlyHandler).Methods("POST")
}

func GetPieByDateHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	dateRange := core.CreateDateRange(startDate, endDate)
	response, err := GetPieByDate(dateRange)
	if err != nil {
		log.Printf("GetPieByDateHandler: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	core.WriteJSONOk(w, response)
}

// Aggregate By month
func GetMetricHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	timePeriod := r.URL.Query().Get("time_period")
	category := r.URL.Query().Get("category")
	dateRange := core.CreateDateRange(startDate, endDate)

	response, err := GetMetric(dateRange, category, timePeriod)
	if err != nil {
		log.Printf("GetPieByDateHandler: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	core.WriteJSONOk(w, response)
}

// func GetGraphByDateHandler(w http.ResponseWriter, r *http.Request) {
// 	response, err := GetMonthlyTransactions()
// 	if err != nil {
// 		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	core.WriteJSONOk(w, response)
// }

// func GetPieMonthlyHandler(w http.ResponseWriter, r *http.Request) {
// 	response, err := GetMonthlyTransactionSumsByCategory(r.URL.Query().Get("name"))
// 	if err != nil {
// 		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	core.WriteJSONOk(w, response)
// }

// func GetGraphMonthlyHandler(w http.ResponseWriter, r *http.Request) {
// 	response, err := GetCategoryAnalysis(r.URL.Query().Get("name"))
// 	if err != nil {
// 		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	core.WriteJSONOk(w, response)
// }

// func GetAnalysisHandler(w http.ResponseWriter, r *http.Request) {
// 	// var dates DateParams
// 	// if err := json.NewDecoder(r.Body).Decode(&dates); err != nil {
// 	// 	log.Printf("GetAnalysisHandler: Error decoding transaction: %v", err)
// 	// 	http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }

// 	response, err := GetCategorySums()
// 	if err != nil {
// 		log.Printf("AddCategoriesHandler: Error adding transaction: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	core.WriteJSONOk(w, response)
// }
