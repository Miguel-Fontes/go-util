package time

import "time"

// IsSameDay is a predicate that verifies if two time.Time types
// represent the same cronological day, comparing day, month and year.
func IsSameDay(time1, time2 time.Time) bool {
	return time1.Day() == time2.Day() &&
		time1.Month() == time2.Month() &&
		time1.Year() == time2.Year()
}
