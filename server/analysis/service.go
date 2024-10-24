package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/shahar05/cash-flow-viewer/categories"
	"github.com/shahar05/cash-flow-viewer/transactions"
)

func GetCategoryAnalysis(categoryName string) ([]transactions.Transaction, error) {

	query := `
    SELECT t.id, t.external_id, t.name, t.amount, t.date, t.address, 
           t.card_unique_id, t.category_id, t.merchant_phone_no, t.international_branch_id
    FROM transactions t
    JOIN categories c ON t.category_id = c.id
    WHERE c.name = $1;
    `

	// Prepare to store the result
	var transArray []transactions.Transaction

	// Execute the query
	rows, err := DB.Query(query, categoryName)
	if err != nil {
		log.Printf("GetTransactionsByCategory: error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		transaction := transactions.Transaction{Category: &categories.Category{}}
		err := rows.Scan(
			&transaction.ID,
			&transaction.ExternalID,
			&transaction.Name,
			&transaction.Amount,
			&transaction.Date,
			&transaction.Address,
			&transaction.CardUniqueId,
			&transaction.Category.ID,
			&transaction.MerchantPhoneNo,
			&transaction.InternationalBranchID,
		)
		if err != nil {
			log.Printf("GetTransactionsByCategory: error scanning row: %v", err)
			return nil, err
		}
		transArray = append(transArray, transaction)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		log.Printf("GetTransactionsByCategory: error iterating rows: %v", err)
		return nil, err
	}

	return transArray, nil
}

// SQL query to get the sum by category within the date range
// query := `
// 	SELECT c.name, SUM(t.amount)
// 	FROM transactions t
// 	JOIN categories c ON t.category_id = c.id
// 	WHERE t.date BETWEEN $1 AND $2
// 	GROUP BY c.name
// `

// Function to query the database and return the results
func GetCategorySums() ([]CategorySum, error) {
	query := `
	SELECT c.name, SUM(t.amount)
	FROM transactions t
	JOIN categories c ON t.category_id = c.id
	GROUP BY c.name
`

	rows, err := DB.Query(query) // , params.StartDate, params.EndDate
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var response []CategorySum

	// Iterate over the result set
	for rows.Next() {
		var categorySum CategorySum
		if err := rows.Scan(&categorySum.Name, &categorySum.Sum); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		response = append(response, categorySum)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return response, nil
}

// MonthlyCategoryTransaction represents the structure for a monthly transaction summary by category
type MonthlyCategoryTransaction struct {
	Month time.Time // Month is the start date of each month
	Total float64   // Total is the sum of transactions in that month for the category
}

// GetMonthlyTransactionSumsByCategory retrieves the total transaction amount per month for a specific category over the past year
func GetMonthlyTransactionSumsByCategory(categoryName string) ([]MonthlyCategoryTransaction, error) {
	// Get the current time and calculate the date range (last year)
	now := time.Now()
	startDate := now.AddDate(-1, 0, 0) // One year ago from today

	// SQL query to get the sum of transactions per month for the specified category
	query := `
        SELECT 
            DATE_TRUNC('month', t.date) AS month, 
            SUM(t.amount) AS total 
        FROM transactions t
        JOIN categories c ON t.category_id = c.id
        WHERE t.date >= $1 AND t.date <= $2
        AND c.name = $3
        GROUP BY month
        ORDER BY month;
    `

	// Execute the query
	rows, err := DB.Query(query, startDate, now, categoryName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the results
	var monthlyCategoryTransactions []MonthlyCategoryTransaction

	// Iterate over the result rows
	for rows.Next() {
		var month time.Time
		var total float64

		// Scan the row into variables
		if err := rows.Scan(&month, &total); err != nil {
			return nil, err
		}

		// Append the result to the slice
		monthlyCategoryTransactions = append(monthlyCategoryTransactions, MonthlyCategoryTransaction{
			Month: month,
			Total: total,
		})
	}

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return monthlyCategoryTransactions, nil
}

// MonthlyTransaction struct to hold monthly transaction data
type MonthlyTransaction struct {
	Month time.Time // Month is the start date of each month
	Total float64   // Total is the sum of transactions in that month
}

// GetMonthlyTransactions function retrieves monthly transaction totals
func GetMonthlyTransactions() ([]MonthlyTransaction, error) {
	query := `
		SELECT DATE_TRUNC('month', date) AS month,
		       SUM(amount) AS total
		FROM transactions
		GROUP BY month
		ORDER BY month;`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var monthlyTransactions []MonthlyTransaction

	for rows.Next() {
		var mt MonthlyTransaction
		if err := rows.Scan(&mt.Month, &mt.Total); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		monthlyTransactions = append(monthlyTransactions, mt)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return monthlyTransactions, nil
}
