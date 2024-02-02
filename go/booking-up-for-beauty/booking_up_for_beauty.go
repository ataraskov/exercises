package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t.Hour() >= 12 && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout = "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.", Schedule(date).Format(layout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	opened := time.Date(2012, 9, 15, 0, 0, 0, 0, time.UTC)
	return time.Date(time.Now().Year(), opened.Month(), opened.Day(), 0, 0, 0, 0, time.UTC)
}
