package core

import (
	"log"
	"time"
)

// ConvertStringToTime converts a time string in multiple formats to time.Time.
// Supported formats:
// - "2024-09-27T22:16:58" (ISO 8601 with time)
// - "2024-09-27" (ISO 8601 date only)
// - "2024/09/27" (Date with slashes)
func ConvertStringToTime(timeStr string) *time.Time {
	formats := []string{
		"2006-01-02T15:04:05", // Full date-time format
		"2006-01-02",          // Date-only format with dashes
		"2006/01/02",          // Date-only format with slashes
	}

	for _, layout := range formats {
		t, err := time.Parse(layout, timeStr)
		if err == nil {
			return &t
		}
	}

	log.Printf("ConvertStringToTime - Error: input '%s' does not match supported formats", timeStr)
	return nil
}

func CreateDateRange(start, end string) *DateRange {
	// Convert start and end strings to time.Time
	startTime := ConvertStringToTime(start)
	endTime := ConvertStringToTime(end)

	// Check if the conversion failed for either date
	if startTime == nil || endTime == nil {
		return nil
	}

	// Validate that start is not after end
	if startTime.After(*endTime) {
		return nil
	}

	return &DateRange{
		StartDate: *startTime,
		EndDate:   *endTime,
	}
}
