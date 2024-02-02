package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First = WeekSchedule(iota) + 1
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	for ; t.Month() == month; t = t.Add(time.Hour * 24) {
		if t.Weekday() != wDay {
			continue
		}

		day := t.Day()
		lastDay := t.AddDate(0, 1, -t.Day()).Day()
		switch {
		case wSched < 5 && day >= int(wSched)*7-6 && day <= int(wSched)*7:
			return day
		case wSched == Teenth && day >= 13 && day <= 19:
			return day
		case wSched == Last && day > lastDay-7 && day <= lastDay:
			return day
		}
	}
	return t.Day()
}
