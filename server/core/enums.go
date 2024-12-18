package core

import "fmt"

// TimePeriod Enum

type TimePeriod string

const (
	Second TimePeriod = "second"
	Minute TimePeriod = "minute"
	Hour   TimePeriod = "hour"
	Day    TimePeriod = "day"
	Week   TimePeriod = "week"
	Month  TimePeriod = "month"
	Year   TimePeriod = "year"
)

func ValidateTimePeriod(timePeriodStr string) (*TimePeriod, error) {
	// Map of valid TimePeriod strings to TimePeriod constants
	validPeriods := map[string]TimePeriod{
		"second": Second,
		"minute": Minute,
		"hour":   Hour,
		"day":    Day,
		"week":   Week,
		"month":  Month,
		"year":   Year,
	}

	// Check if the timePeriodStr is a valid key in the map
	if period, valid := validPeriods[timePeriodStr]; valid {
		return &period, nil
	}

	// Print the error and return nil if invalid
	err := fmt.Errorf("invalid time period: %s", timePeriodStr)
	fmt.Println(err) // print the error to the console
	return nil, err
}
