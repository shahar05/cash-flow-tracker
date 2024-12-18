package core

import "time"

type ResponseStructure struct {
	Data   interface{} `json:"data,omitempty"`
	Error  *ErrorInfo  `json:"error,omitempty"`
	Status bool        `json:"status"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ColumnValue struct to hold column names and their corresponding values
type ColumnValue struct {
	ColName string
	Value   interface{}
}

// Input struct for date parameters
type DateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Metric represents a metric with a specific time period and a list of aggregated values.
type Metric struct {
	Period         TimePeriod        `json:"time_period"`     // Changed Enum to a named type for clarity
	AggregatedData []AggregatedValue `json:"aggregated_data"` // Renamed result to clarify its purpose
}

// AggregatedValue holds the time and value for an aggregated metric.
type AggregatedValue struct {
	Timestamp time.Time `json:"timestamp"` // Renamed time to Timestamp for clarity
	Value     float64   `json:"value"`
}
