package utils

import (
	"time"
)

func NormalizeDate(dateStr string) (time.Time, string, error) {
	var t time.Time
	var err error

	if dateStr == "" {
		t = time.Now()
	} else {
		t, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return time.Time{}, "", err
		}
	}

	standardized := t.UTC().Truncate(time.Second)

	return standardized, standardized.Format(time.RFC3339), nil
}
