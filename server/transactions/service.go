package transactions

import "log"

// AddTransaction adds a new transaction to the database and returns the ID
func AddTransaction(transaction Transaction) (string, error) {
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
