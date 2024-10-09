package transactions

import (
	"time"

	"github.com/shahar05/cash-flow-viewer/categories"
)

type Transaction struct {
	ID                    string     `json:"id"`                          // Primary Key (Auto Generated)
	Name                  string     `json:"name"`                        // Not null
	Amount                float64    `json:"amount"`                      // Not null
	DateStr               string     `json:"date_str"`                    // Not in SQL
	Date                  *time.Time `json:"date"`                        // Not null
	Address               *string    `json:"address,omitempty"`           // Can be null
	CardUniqueId          *string    `json:"card_unique_id,omitempty"`    // Can be null
	Category              string     `json:"category"`                    // Foreign Key Category Name
	MerchantPhoneNo       *string    `json:"merchant_phone_no,omitempty"` // Can be null
	InternationalBranchID string     `json:"international_branch_id"`     // Not null
}

type AttachTransReq struct {
	Trans    *Transaction         `json:"transaction"`
	Category *categories.Category `json:"category"`
}

/*
   {
       "merchantName": "WOLT",
       "amountForDisplay": 136,
       "debCrdCurrencyDesc": "שח",
       "debCrdDate": "2024-05-02T00:00:00",
       "merchantAddress": "יבנה 40 תל אביב - יפו",
       "merchantPhoneNo": "03-7631092",
       "merchantId": "958376",
       "cardUniqueId": "338994664018201955",
       "internationalBranchID": "5411"
   }
*/
