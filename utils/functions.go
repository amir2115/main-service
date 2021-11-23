package utils

import "time"

func GetCurrentTime() time.Time {
	utc := (3 * time.Hour) + (30 * time.Minute)
	return time.Now().Add(utc)
}
