package utils

import (
	"log"
	"math/rand"
	"time"
)

func Ptr[T any](v T) *T {
	return &v
}

// GetRandomInRange generates a random integer between min and max (inclusive)
func GetRandomInRange(min, max int) int {
	if min > max {
		log.Printf("Invalid range: min (%d) is greater than max (%d)\n", min, max)
		return -1 // -1 equal to error
	}
	return rand.Intn(max-min+1) + min
}

// ConvertStringToTime converts a time string in "2024-09-27T22:16:58" format to time.Time
func ConvertStringToTime(timeStr string) *time.Time {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Printf("ConvertStringToTime: Error: %s", err.Error())
		return nil
	}
	return &t
}
