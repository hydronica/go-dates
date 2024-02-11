package dates

import (
	"log/slog"
	"time"
)

// Default start and end dates for the week
// can be changed if the definition of a week is different
// for example in the US, the week starts on Sunday

// a helper library for date (no time) calculations
// for example, getting the start and end dates of the current week, previous week, etc.

const (
	OneDay       = time.Hour * 24 // duration for one day 24 hours
	OneWeek      = OneDay * 7     // duration for one week 7 days or 168 hours
	StartDefault = time.Monday    // default weekday start of the week
	EndDefault   = time.Sunday    // default weekday end of the week
)

type Week struct {
	weekStart time.Weekday // starting weekday of the week
	weekEnd   time.Weekday // ending weekday of the week
}

// New returns a new Week with the given start and end days.
// First value should be start of week, second value should be end of week.
// For example New(time.Monday, time.Sunday), extra values are ignored
// leave empty to use the default week start and end weekdays i.e., NewWeek()
func NewWeek(day ...time.Weekday) Week {
	var day1, day2 time.Weekday

	for i, d := range day {
		switch i {
		case 0:
			day1 = d
		case 1:
			day2 = d
		}
	}

	w := Week{
		weekStart: StartDefault,
		weekEnd:   EndDefault,
	}

	if len(day) < 2 {
		slog.Warn("not enough days given, using default")
		return w
	}

	if day1 == time.Sunday && day1 != time.Saturday {
		slog.Warn("there are not 7 days in given week, using default")
		return w
	}
	test := day1 - 1
	if day1 > time.Sunday && test != day2 {
		slog.Warn("week start and end days are not consecutive, using default")
		return w
	}

	w.weekStart = day1
	w.weekEnd = day2

	return w
}

// Date provides a shorthand for createing a time.Date truncating the time.
// timezone is ignored and UTC is used
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Day returns the tuncated date of the given time t
func Day(t time.Time) time.Time {
	return Date(t.Year(), t.Month(), t.Day())
}

// LastDayOfMonth from the given t time
func LastDayOfMonth(t time.Time) time.Time {
	return Date(t.Year(), t.Month()+1, 0)
}

// WeekAdd returns t time with weeks added (use negative value to subtract)
func WeekAdd(t time.Time, weeks int) time.Time {
	t = t.Add(OneWeek * time.Duration(weeks))
	return t
}

// PriorLastFullWeek returns the start and end dates of the week prior to the last full week
// or two weeks ago
func (d Week) PriorLastFullWeek(t time.Time) (start, end time.Time) {
	lfwStart, _ := d.LastFullWeek(t)
	start = lfwStart.Add(-OneWeek)
	end = start.Add(OneDay * 6)
	return start, end
}

// StartOfWeek reutrns the date of the of the start of the week less than or equal to the given date t,
// which is the first day of the week back from the given time t
func (d Week) StartOfWeek(t time.Time) time.Time {
	// time is already the start of the week
	if t.Weekday() == d.weekStart {
		return Date(t.Year(), t.Month(), t.Day())
	}
	// subtract days until we reach the start of the week
	for t.Weekday() != d.weekStart {
		t = t.Add(-OneDay)
	}
	return Date(t.Year(), t.Month(), t.Day())
}

// LastFullWeek returns the start and end dates of the last full week
func (d Week) LastFullWeek(t time.Time) (start, end time.Time) {
	t = d.StartOfWeek(t)
	start = t.Add(-OneWeek)
	end = start.Add(OneDay * 6)
	return start, end
}

// PrevYearLastFullWeek returns the start and end dates of the last full week of the previous year
func (d Week) PrevYearLastFullWeek(t time.Time) (start, end time.Time) {
	startLfw, _ := d.LastFullWeek(t)
	start = d.StartOfWeek(startLfw.Add(OneDay * -363)) // this is attempting to account for leap year
	end = start.Add(OneDay * 6)
	return start, end
}

// MonthToDate returns the start and end dates of the current month
func MonthToDate(t time.Time) (start, end time.Time) {
	start = Date(t.Year(), t.Month(), 1)
	end = t
	return start, end
}

// FullMonth returns the start and end dates of the current month
func FullMonth(t time.Time) (start, end time.Time) {
	start = Date(t.Year(), t.Month(), 1)
	end = FirstOfNextMonth(t).Add(-OneDay)
	return start, end
}

// FirstOfNextMonth returns the 1st of the next month from time t
func FirstOfNextMonth(t time.Time) time.Time {
	year := t.Year()
	nextMonth := t.Month() + 1
	if nextMonth > time.December {
		nextMonth = time.January
		year = t.Year() + 1
	}
	return Date(year, nextMonth, 1)
}

// PrevMonth returns the start and end dates of the previous month
func PrevMonth(t time.Time) (start, end time.Time) {
	prevMonth := t.Month() - 1
	year := t.Year()
	if prevMonth < 1 {
		prevMonth = time.December
		year = t.Year() - 1
	}
	t = Date(year, prevMonth, t.Day())

	start = Date(year, prevMonth, 1)
	end = FirstOfNextMonth(t).Add(-OneDay)

	return start, end
}

// PrevMonthToDate returns the start and end dates of the previous month to the given date (t)
func PrevMonthToDate(t time.Time) (start, end time.Time) {
	prevMonth := t.Month() - 1
	year := t.Year()
	if prevMonth < 1 {
		prevMonth = time.December
		year = t.Year() - 1
	}
	lm := Date(year, prevMonth, t.Day())
	// if subtracking a month results in a date in the same month as t
	// this means the day is greater than the last day of the previous month
	if lm.Month() == t.Month() {
		return PrevMonth(lm)
	}

	return MonthToDate(lm)
}

// PrevYearMtd returns the start and end dates up to t
// of the same month in the previous year
// if a leap day is given for t the previous year's last day will be feb 28th
func PrevYearMtd(t time.Time) (start, end time.Time) {
	prevYear := t.Year() - 1

	pEnd := Date(prevYear, t.Month(), t.Day())
	if pEnd.Month() > t.Month() {
		pEnd = pEnd.Add(-OneDay)
	}

	start = StartOfMonth(pEnd)

	return start, pEnd
}

// YearToDate returns the start and end dates of the current year
func YearToDate(t time.Time) (start, end time.Time) {
	start = Date(t.Year(), 1, 1)
	end = t
	return start, end
}

// PreviousYearToDate returns the start and end dates of the previous year
func PreviousYearToDate(t time.Time) (start, end time.Time) {
	start = Date(t.Year()-1, 1, 1)
	end = Date(t.Year()-1, t.Month(), t.Day())
	// accounts for leap day
	if end.Month() > t.Month() {
		end = end.Add(-OneDay)
	}
	return start, end
}

// StartOfMonth returns 1st of current month @ midnight UTC
func StartOfMonth(t time.Time) time.Time {
	return Date(t.Year(), t.Month(), 1)
}
