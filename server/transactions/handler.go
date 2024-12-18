package transactions

import (
	"database/sql"
	"encoding/json"
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
	r.HandleFunc("/transactions", GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions", AddTransactionsHandler).Methods("POST")
	r.HandleFunc("/attach-transaction", AttachTransactionHandler).Methods("POST")
	r.HandleFunc("/filter-transactions", FilterTransactionsHandler).Methods("POST")
}

func FilterTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("FilterTransactionsHandler: Handling POST /filter-transactions request")
	var body FullTransDetailArray
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("AttachTransactionHandler: Error decoding AttachTransReq: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filteredTrans, err := FilterTransactions(body.FullTransArr)
	if err != nil {
		core.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	core.WriteJSONOk(w, filteredTrans)
}

// GetTransactionsHandler handles GET requests for contacts with pagination
func AttachTransactionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AttachTransactionHandler: Handling POST /transactions request")
	var attachTransReq AttachTransReq
	if err := json.NewDecoder(r.Body).Decode(&attachTransReq); err != nil {
		log.Printf("AttachTransactionHandler: Error decoding AttachTransReq: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

// GetTransactionsHandler handles GET requests for contacts with pagination
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {

}

// AddTransactionsHandler handles POST requests to add a new transaction
func AddTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AddTransactionsHandler: Handling POST /transactions request")
	var trans Transaction
	if err := json.NewDecoder(r.Body).Decode(&trans); err != nil {
		log.Printf("AddTransactionsHandler: Error decoding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := AddTransaction(trans)
	if err != nil {
		log.Printf("AddTransactionsHandler: Error adding transaction: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	trans.ID = id
	log.Printf("AddTransactionsHandler: Added new transaction with ID: %s", id)
	core.WriteJSONOk(w, trans)
}
