package dates

import (
	"time"
)

// New Years Day
func NewYearsDay(date time.Time) time.Time {
	return Date(date.Year(), time.January, 1)
}

func NewYearsEve(date time.Time) time.Time {
	return Date(date.Year(), time.December, 31)
}

// Martin Luther King Jr. Day
func MartinLutherKingJrDay(date time.Time) time.Time {
	// first of the given year
	date = Date(date.Year(), time.January, 1)
	count := 0
	for {
		if date.Weekday() == time.Monday {
			count++
		}
		if count >= 3 {
			break
		}
		date = date.Add(OneDay)
	}
	return date
}

func MemorialDay(date time.Time) time.Time {
	// first of the given year
	date = Date(date.Year(), time.June, 1)
	for {
		if date.Weekday() == time.Monday {
			break
		}
		date = date.Add(-OneDay)
	}
	return date
}

func Juneteenth(date time.Time) time.Time {
	return Date(date.Year(), time.June, 19)
}

func IndependenceDay(date time.Time) time.Time {
	return Date(date.Year(), time.July, 4)
}

func LaborDay(date time.Time) time.Time {
	date = Date(date.Year(), time.September, 1)
	for {
		if date.Weekday() == time.Monday {
			break
		}
		date = date.Add(OneDay)
	}
	return date
}

func VeteransDay(date time.Time) time.Time {
	return Date(date.Year(), time.November, 11)
}
