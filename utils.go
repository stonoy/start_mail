package main

import (
	"fmt"
	"time"
)

func getTimeFromStr(timeStr string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04",
	}

	for _, layout := range layouts {
		theTime, err := time.Parse(layout, timeStr)
		if err == nil {
			return theTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("error in parsing time %v", timeStr)
}
