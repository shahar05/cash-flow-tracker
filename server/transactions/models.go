package transactions

import (
	"time"

	"github.com/shahar05/cash-flow-viewer/categories"
)

type Transaction struct {
	ID                    string               `json:"id"`                          // Primary Key (Auto Generated)
	Name                  string               `json:"name"`                        // Not null
	Amount                float64              `json:"amount"`                      // Not null
	DateStr               string               `json:"date_str"`                    // Not in SQL
	Date                  *time.Time           `json:"date"`                        // Not null
	Address               *string              `json:"address,omitempty"`           // Can be null
	CardUniqueId          *string              `json:"card_unique_id,omitempty"`    // Can be null
	Category              *categories.Category `json:"category"`                    // Foreign Key Category Name
	MerchantPhoneNo       *string              `json:"merchant_phone_no,omitempty"` // Can be null
	InternationalBranchID string               `json:"international_branch_id"`     // Not null
}

type AttachTransReq struct {
	Trans    *Transaction         `json:"transaction"`
	Category *categories.Category `json:"category"`
}

type FullTransDetail struct {
	PaymentsMade                []interface{} `json:"paymentsMade"`
	PaymentsMadeComment         *string       `json:"paymentsMadeComment,omitempty"`
	EarlyPaymentInd             bool          `json:"earlyPaymentInd"`
	TrnIntId                    string        `json:"trnIntId"`
	CardUniqueId                string        `json:"cardUniqueId"`
	MerchantName                string        `json:"merchantName"`
	AmountForDisplay            float64       `json:"amountForDisplay"`
	CurrencyForDisplay          string        `json:"currencyForDisplay"`
	TrnPurchaseDate             string        `json:"trnPurchaseDate"`
	TrnAmt                      float64       `json:"trnAmt"`
	TrnCurrencyIsoCode          *string       `json:"trnCurrencyIsoCode,omitempty"`
	TrnCurrencySymbol           string        `json:"trnCurrencySymbol"`
	DebCrdCurrencyDesc          string        `json:"debCrdCurrencyDesc"`
	DebCrdCurrencyCode          int           `json:"debCrdCurrencyCode"`
	DebCrdDate                  string        `json:"debCrdDate"`
	AmtBeforeConvAndIndex       float64       `json:"amtBeforeConvAndIndex"`
	DebCrdCurrencySymbol        string        `json:"debCrdCurrencySymbol"`
	MerchantAddress             string        `json:"merchantAddress"`
	MerchantPhoneNo             string        `json:"merchantPhoneNo"`
	BranchCodeDesc              *string       `json:"branchCodeDesc,omitempty"`
	TransCardPresentInd         bool          `json:"transCardPresentInd"`
	CurPaymentNum               int           `json:"curPaymentNum"`
	NumOfPayments               int           `json:"numOfPayments"`
	TrnType                     string        `json:"trnType"`
	TransMahut                  string        `json:"transMahut"`
	TrnNumaretor                int           `json:"trnNumaretor"`
	Comments                    []interface{} `json:"comments"`
	LinkedComments              []interface{} `json:"linkedComments"`
	TokenInd                    int           `json:"tokenInd"`
	WalletProviderCode          int           `json:"walletProviderCode"`
	WalletProviderDesc          string        `json:"walletProviderDesc"`
	TokenNumberPart4            string        `json:"tokenNumberPart4"`
	RoundingAmount              *float64      `json:"roundingAmount,omitempty"`
	RoundingReason              *string       `json:"roundingReason,omitempty"`
	DiscountAmount              *float64      `json:"discountAmount,omitempty"`
	DiscountReason              *string       `json:"discountReason,omitempty"`
	InternationalBranchID       string        `json:"internationalBranchID"`
	TransTypeCommentDetails     []interface{} `json:"transTypeCommentDetails"`
	InternationalBranchDesc     string        `json:"internationalBranchDesc"`
	ChargeExternalToCardComment string        `json:"chargeExternalToCardComment"`
	SuperBranchDesc             *string       `json:"superBranchDesc,omitempty"`
	TransactionTypeCode         int           `json:"transactionTypeCode"`
	RefundInd                   bool          `json:"refundInd"`
	IsImmediate                 bool          `json:"isImmediate"`
	IsImmediateCommentInd       bool          `json:"isImmediateCommentInd"`
	IsImmediateHHKInd           bool          `json:"isImmediateHHKInd"`
	ImmediateComments           []interface{} `json:"immediateComments"`
	IsMargaritaInd              bool          `json:"isMargaritaInd"`
	IsSpreadPaymenstAbroadInd   bool          `json:"isSpreadPaymenstAbroadInd"`
	TrnExacWay                  int           `json:"trnExacWay"`
	IsInterestTransaction       bool          `json:"isInterestTransaction"`
	OnGoingTransactionsComment  string        `json:"onGoingTransactionsComment"`
	MerchantId                  string        `json:"merchantId"`
	CrdExtIdNumTypeCode         string        `json:"crdExtIdNumTypeCode"`
	TransSource                 string        `json:"transSource"`
	TransIndication             string        `json:"transIndication"`
	CashAccountTrnAmt           float64       `json:"cashAccountTrnAmt"`
	TransOriginalSumCurrency    string        `json:"transOriginalSumCurrency"`
	CrmIccCurrencyDesc          string        `json:"crmIccCurrencyDesc"`
	TransOriginalCurrencyCode   *string       `json:"transOriginalCurrencyCode,omitempty"`
	IsAbroadTransaction         bool          `json:"isAbroadTransaction"`
	TransSafeIndication         string        `json:"transSafeIndication"`
	PosEntryMod                 string        `json:"posEntryMod"`
	WalletTokenInd              *int          `json:"walletTokenInd,omitempty"`
	SrcCurrencyCode             *string       `json:"srcCurrencyCode,omitempty"`
	TokenNumber                 string        `json:"tokenNumber"`
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
