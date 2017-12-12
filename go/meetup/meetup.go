package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	switch week {
	case First, Second, Third, Fourth:
		t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		wday := t.Weekday()
		weekdaysDiff := (weekday - wday + 7) % 7
		return int(week)*7 + 1 + int(weekdaysDiff)
	case Teenth:
		t := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		wday := t.Weekday()
		weekdaysDiff := (weekday - wday + 7) % 7
		return 13 + int(weekdaysDiff)
	case Last:
		t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		wday := t.Weekday()
		weekdaysDiff := weekday - wday
		if weekday > wday {
			return t.Day() + int(weekdaysDiff) - 7
		}
		return t.Day() + int(weekdaysDiff)
	default:
		panic("invalid date")
	}
}
