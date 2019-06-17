package main

import "time"

// Timezone represents the used timezone
var Timezone = time.UTC

// DateToTime is a helper function for converting a DMY time to a time.Time object
func DateToTime(day int, month time.Month, year int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, Timezone)
}
