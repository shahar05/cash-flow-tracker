package categories

type CategoryArray struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"` // Primary Key
}

// *Can also work like this: InternationalBranchIDs []string

type CategoryLink struct {
	InternationalBranchID string `json:"international_branch_id"` // Primary Key
	CategoryName          string `json:"category_name"`           // Foreign Key Category ID
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
