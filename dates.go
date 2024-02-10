package dates

import "time"

// Default start and end dates for the week
// can be changed if the definition of a week is different
// for example in the US, the week starts on Sunday

// a helper library for date (no time) calculations

type Base struct {
	weekStart time.Weekday  // starting weekday of the week
	weekEnd   time.Weekday  // ending weekday of the week
	oneDay    time.Duration // default 24 hours
}

func NewBase(weekStart, weekEnd time.Weekday) Base {
	return Base{
		weekStart: weekStart,
		weekEnd:   weekEnd,
		oneDay:    time.Hour * 24,
	}
}

// Date provides a shorthand for createing a time.Date truncating the time.
// timezone is ignored and UTC is used
func (d Base) Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// Day returns the tuncated date of the given time t
func (d Base) Day(t time.Time) time.Time {
	return d.Date(t.Year(), t.Month(), t.Day())
}

// LastDayOfMonth from the given t time
func (d Base) LastDayOfMonth(t time.Time) time.Time {
	return d.Date(t.Year(), t.Month()+1, 0)
}

// WeekAdd returns t time with weeks added (use negative value to subtract)
func (d Base) WeekAdd(t time.Time, weeks int) time.Time {
	t = t.Add(d.oneDay * 7 * time.Duration(weeks))
	return t
}

// StartOfWeek reutrns the date of the of the start of the week less than or equal to the given date
func (d Base) StartOfWeek(t time.Time) time.Time {
	// time is already the start of the week
	if t.Weekday() == d.weekStart {
		return d.Date(t.Year(), t.Month(), t.Day())
	}
	// subtract days until we reach the start of the week
	for t.Weekday() != d.weekStart {
		t = t.Add(-d.oneDay)
	}
	return d.Date(t.Year(), t.Month(), t.Day())
}

// LastFullWeek returns the start and end dates of the last full week
func (d Base) LastFullWeek(t time.Time) (start, end time.Time) {
	t = d.StartOfWeek(t)
	start = t.Add(d.oneDay * -7)
	end = start.Add(d.oneDay * 6)
	return start, end
}

// PriorLastFullWeek returns the start and end dates of the week prior to the last full week
func (d Base) PriorLastFullWeek(t time.Time) (start, end time.Time) {
	lfwStart, _ := d.LastFullWeek(t)
	start = lfwStart.Add(d.oneDay * -7)
	end = start.Add(d.oneDay * 6)
	return start, end
}

// PrevYearLastFullWeek returns the start and end dates of the last full week of the previous year
func (d Base) PrevYearLastFullWeek(t time.Time) (start, end time.Time) {
	startLfw, _ := d.LastFullWeek(t)
	start = d.StartOfWeek(startLfw.Add(d.oneDay * -363))
	end = start.Add(d.oneDay * 6)
	return start, end
}

// MonthToDate returns the start and end dates of the current month
func (d Base) MonthToDate(t time.Time) (start, end time.Time) {
	start = d.Date(t.Year(), t.Month(), 1)
	end = t
	return start, end
}

// FullMonth returns the start and end dates of the current month
func (d Base) FullMonth(t time.Time) (start, end time.Time) {
	start = d.Date(t.Year(), t.Month(), 1)
	end = d.FirstOfNextMonth(t).Add(-d.oneDay)
	return start, end
}

// FirstOfNextMonth returns the 1st of the next month from time t
func (d Base) FirstOfNextMonth(t time.Time) time.Time {
	year := t.Year()
	nextMonth := t.Month() + 1
	if nextMonth > time.December {
		nextMonth = time.January
		year = t.Year() + 1
	}
	return d.Date(year, nextMonth, 1)
}

// PrevMonth returns the start and end dates of the previous month
func (d Base) PrevMonth(t time.Time) (start, end time.Time) {
	prevMonth := t.Month() - 1
	year := t.Year()
	if prevMonth < 1 {
		prevMonth = time.December
		year = t.Year() - 1
	}
	t = d.Date(year, prevMonth, t.Day())

	start = d.Date(year, prevMonth, 1)
	end = d.FirstOfNextMonth(t).Add(-d.oneDay)

	return start, end
}

// PrevMonthToDate returns the start and end dates of the previous month to the given date (t)
func (d Base) PrevMonthToDate(t time.Time) (start, end time.Time) {
	prevMonth := t.Month() - 1
	year := t.Year()
	if prevMonth < 1 {
		prevMonth = time.December
		year = t.Year() - 1
	}
	lm := d.Date(year, prevMonth, t.Day())
	// if subtracking a month results in a date in the same month as t
	// this means the day is greater than the last day of the previous month
	if lm.Month() == t.Month() {
		return d.PrevMonth(lm)
	}

	return d.MonthToDate(lm)
}

// PrevYearMtd returns the start and end dates up to t
// of the same month in the previous year
// if a leap day is given for t the previous year's last day will be feb 28th
func (d Base) PrevYearMtd(t time.Time) (start, end time.Time) {
	prevYear := t.Year() - 1

	pEnd := d.Date(prevYear, t.Month(), t.Day())
	if pEnd.Month() > t.Month() {
		pEnd = pEnd.Add(-d.oneDay)
	}

	start = d.StartOfMonth(pEnd)

	return start, pEnd
}

// YearToDate returns the start and end dates of the current year
func (d Base) YearToDate(t time.Time) (start, end time.Time) {
	start = d.Date(t.Year(), 1, 1)
	end = t
	return start, end
}

// PreviousYearToDate returns the start and end dates of the previous year
func (d Base) PreviousYearToDate(t time.Time) (start, end time.Time) {
	start = d.Date(t.Year()-1, 1, 1)
	end = d.Date(t.Year()-1, t.Month(), t.Day())
	// accounts for leap day
	if end.Month() > t.Month() {
		end = end.Add(-d.oneDay)
	}
	return start, end
}

// StartOfMonth returns 1st of current month @ midnight UTC
func (d Base) StartOfMonth(t time.Time) time.Time {
	return d.Date(t.Year(), t.Month(), 1)
}
