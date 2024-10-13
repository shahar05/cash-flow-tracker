package transactions

import "log"

func FilterTransactions([]FullTransDetail) ([]FullTransDetail, error) {
	// Prepare SQL query
	query := `
		SELECT id, name, amount, date, address, card_unique_id, category_id, merchant_phone_no, international_branch_id
		FROM filter_transactions($1, $2, $3);`

	// Execute the query
	rows, err := db.Query(query, name, amount, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []FullTransDetail

	// Iterate through the rows
	for rows.Next() {
		var transaction FullTransDetail
		if err := rows.Scan(&transaction.ID, &transaction.Name, &transaction.Amount, &transaction.Date, &transaction.Address, &transaction.CardUniqueId, &transaction.CategoryId, &transaction.MerchantPhoneNo, &transaction.InternationalBranchID); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

// AddTransaction adds a new transaction to the database and returns the ID
func AddTransaction(transaction Transaction) (string, error) {
	log.Printf("AddTransaction: Adding transaction %v", transaction)
	var id string
	err := DB.QueryRow(
		`INSERT INTO transactions 
		(name, amount, date, address, card_unique_id, category_id, merchant_phone_no, international_branch_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id`,
		transaction.Name, transaction.Amount, transaction.Date, transaction.Address, transaction.CardUniqueId,
		transaction.Category.ID, transaction.MerchantPhoneNo, transaction.InternationalBranchID,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func AttachTransaction(transaction Transaction) (string, error) {
	log.Printf("AddTransaction: Adding transaction %v", transaction)
	var id string
	err := DB.QueryRow(
		`INSERT INTO transactions 
		(name, amount, date, address, card_unique_id, category, merchant_phone_no, international_branch_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id`,
		transaction.Name, transaction.Amount, transaction.Date, transaction.Address, transaction.CardUniqueId,
		transaction.Category, transaction.MerchantPhoneNo, transaction.InternationalBranchID,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
