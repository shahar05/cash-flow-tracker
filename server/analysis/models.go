package analysis

import "time"

// Input struct for date parameters
type DateParams struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Result struct for category and sum
type CategorySum struct {
	Name string  `json:"name"`
	Sum  float64 `json:"sum"`
}

type MonthlyCategorySum struct {
	Month          string        `json:"month"` // YYYY-MM format
	CategorySumArr []CategorySum `json:"categorySum"`
}
