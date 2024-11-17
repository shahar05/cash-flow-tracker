package transactions

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/shahar05/cash-flow-viewer/categories"
	"github.com/shahar05/cash-flow-viewer/utils"
)

func FilterTransactions(fullTransArr []FullTransDetail) ([]FullTransDetail, error) {
	// Step 1: Collect all transaction internal IDs (trnIntId) from the input array
	trnIntIds := make([]string, 0, len(fullTransArr))
	for _, trans := range fullTransArr {
		trnIntIds = append(trnIntIds, trans.TrnIntId)
	}

	// Step 2: Fetch existing transactions in bulk based on trnIntIds
	existingTransMap := make(map[string]struct{})
	if err := fetchExistingTransactions(trnIntIds, existingTransMap); err != nil {
		log.Printf("FilterTransactions: Error fetching existing transactions: %v", err)
		return nil, err
	}

	// Step 3: Filter and insert non-existing transactions
	var nonInsertedTransactions []FullTransDetail
	for _, trans := range fullTransArr {
		// Process only transactions that don't already exist
		if _, exists := existingTransMap[trans.TrnIntId]; !exists {
			inserted := tryToInsertTransaction(trans)
			if !inserted {
				nonInsertedTransactions = append(nonInsertedTransactions, trans)
			}
		}
	}

	return nonInsertedTransactions, nil
}

// AddTransaction adds a new transaction to the database and returns the ID
func AddTransaction(transaction Transaction) (string, error) {
	log.Printf("AddTransaction: Adding transaction %v", transaction)

	// Ensure date is not nil
	if transaction.Date == nil {
		return "", fmt.Errorf("AddTransaction: transaction date is required")
	}

	if transaction.Amount < 0 { //attach to credit Category
		transaction.Category.ID = "2d496035-9ed0-4511-a461-318326f07390"
	}

	var id string
	err := DB.QueryRow(
		`INSERT INTO transactions 
		(external_id, name, amount, date, address, card_unique_id, category_id, merchant_phone_no, international_branch_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id`,
		transaction.ExternalID,            // trnIntId
		transaction.Name,                  // merchantName
		transaction.Amount,                // amountForDisplay
		transaction.Date,                  // trnPurchaseDate
		transaction.Address,               // nullable merchantAddress
		transaction.CardUniqueId,          // nullable CardUniqueId
		transaction.Category.ID,           // category_id (Foreign Key)
		transaction.MerchantPhoneNo,       // nullable MerchantPhoneNo
		transaction.InternationalBranchID, // InternationalBranchID
	).Scan(&id)

	if err != nil {
		log.Printf("AddTransaction: Error inserting transaction: %v", err)
		return "", err
	}

	log.Printf("AddTransaction: Successfully added transaction with ID %s", id)
	return id, nil
}

// tryToInsertTransaction attempts to insert a transaction by auto-assigning fields
func tryToInsertTransaction(fullTrans FullTransDetail) bool {
	trans := Transaction{
		Category: &categories.Category{}, // Initialize Category to avoid nil pointer dereference
	}
	transName := strings.TrimSpace(fullTrans.MerchantName)

	// Search for existing transactions by name or international_branch_id
	query := `  (
    				SELECT category_id 
    				FROM transactions 
    				WHERE name = $1 OR international_branch_id = $2 
    				LIMIT 1
				)
					UNION
				(
				    SELECT category_id 
				    FROM category_links 
				    WHERE international_branch_id = $2 
				    LIMIT 1
			    )
			        LIMIT 1;`

	if err := DB.QueryRow(query, transName, fullTrans.InternationalBranchID).Scan(&trans.Category.ID); err != nil {
		if err == sql.ErrNoRows {
			// No matching transaction found
			log.Printf("tryToInsertTransaction: No matching transaction found for %v.", fullTrans)
			return false
		}
		log.Printf("tryToInsertTransaction: Query error for %v. Error: %v", fullTrans, err)
		return false
	}

	// Create a transaction object from full transaction details
	createTransFromFullTrans(&trans, &fullTrans)

	// Insert the transaction into the database
	if _, err := AddTransaction(trans); err != nil {
		log.Printf("tryToInsertTransaction: Failed to add transaction. Error: %v", err)
		return false
	}

	return true
}

// createTransFromFullTrans maps full transaction details to a Transaction object
func createTransFromFullTrans(t *Transaction, ft *FullTransDetail) {
	t.Address = &ft.MerchantAddress
	t.Amount = ft.AmountForDisplay
	t.CardUniqueId = &ft.CardUniqueId
	t.DateStr = ft.TrnPurchaseDate
	t.Date = utils.ConvertStringToTime(ft.TrnPurchaseDate)
	t.InternationalBranchID = ft.InternationalBranchID
	t.MerchantPhoneNo = &ft.MerchantPhoneNo
	t.Name = ft.MerchantName
	t.ExternalID = ft.TrnIntId
}

// fetchExistingTransactions retrieves transactions from the database based on trnIntIds
func fetchExistingTransactions(trnIntIds []string, existingTransMap map[string]struct{}) error {
	if len(trnIntIds) == 0 {
		return nil
	}

	// Create the SQL query with placeholders for trnIntIds
	placeholders := make([]string, len(trnIntIds))
	for i := range trnIntIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("SELECT external_id FROM transactions WHERE external_id IN (%s)", strings.Join(placeholders, ","))

	// Execute the query
	rows, err := DB.Query(query, convertToInterfaceSlice(trnIntIds)...)
	if err != nil {
		log.Printf("fetchExistingTransactions: Error executing query: %v", err)
		return err
	}
	defer rows.Close()

	// Scan the results and populate existingTransMap
	for rows.Next() {
		var externalId string
		if err := rows.Scan(&externalId); err != nil {
			log.Printf("fetchExistingTransactions: Error scanning row: %v", err)
			return err
		}
		existingTransMap[externalId] = struct{}{}
	}

	return rows.Err()
}

// convertToInterfaceSlice converts a slice of strings to a slice of empty interfaces for SQL query arguments
func convertToInterfaceSlice(slice []string) []interface{} {
	interfaces := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaces[i] = v
	}
	return interfaces
}

// BuildValueString generates a SQL value string for batch inserts with dynamic parameters.
// n represents how many sets of values are needed (e.g., for each transaction).
// paramsPerSet represents how many parameters are there in each set.
func BuildValueString(n, paramsPerSet int) string {
	var valueStrings []string

	for i := 0; i < n; i++ {
		var placeholders []string
		for j := 0; j < paramsPerSet; j++ {
			placeholders = append(placeholders, fmt.Sprintf("$%d", i*paramsPerSet+j+1))
		}
		valueStrings = append(valueStrings, fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")))
	}
	return strings.Join(valueStrings, ", ")
}

// InsertTransactions inserts multiple transactions into the database in a single batch
func InsertTransactions(transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	// List of columns for the query
	cols := []string{"external_id", "name", "amount", "date", "address", "card_unique_id", "category_id", "merchant_phone_no", "international_branch_id"}

	// Build the SQL query with dynamic value placeholders
	query := fmt.Sprintf(`INSERT INTO transactions (%s) VALUES`, strings.Join(cols, ", "))

	// Use BuildValueString to create the placeholders for the number of transactions
	query += BuildValueString(len(transactions), len(cols))

	// Prepare the query to return the inserted transaction IDs
	query += " RETURNING id"

	// Prepare the arguments for the SQL query
	var args []interface{}
	for _, tr := range transactions {
		args = append(args, tr.ExternalID, tr.Name, tr.Amount, tr.Date, tr.Address, tr.CardUniqueId, tr.Category.ID, tr.MerchantPhoneNo, tr.InternationalBranchID)
	}

	// Execute the query
	stmt, err := DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement with arguments
	rows, err := stmt.Query(args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Scan the returned transaction IDs
	for rows.Next() {
		var insertedID string
		if err := rows.Scan(&insertedID); err != nil {
			return fmt.Errorf("failed to scan inserted ID: %w", err)
		}
		fmt.Printf("Inserted transaction with ID: %s\n", insertedID)
	}

	return nil
}
