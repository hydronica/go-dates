package dates

import (
	"testing"
	"time"

	"github.com/hydronica/trial"
)

type output struct {
	start time.Time
	end   time.Time
}

func TestNewWeek(t *testing.T) {

	fn := func(in []time.Weekday) (Week, error) {
		return NewWeek(in...), nil
	}

	cases := trial.Cases[[]time.Weekday, Week]{
		"default": {
			Input:    nil,
			Expected: Week{weekStart: StartDefault, weekEnd: EndDefault},
		},
		"tues to mon": {
			Input:    []time.Weekday{time.Tuesday, time.Monday},
			Expected: Week{weekStart: time.Tuesday, weekEnd: time.Monday},
		},
		"fri to thurs": {
			Input:    []time.Weekday{time.Friday, time.Thursday},
			Expected: Week{weekStart: time.Friday, weekEnd: time.Thursday},
		},
		"invalid week": {
			Input:    []time.Weekday{time.Monday, time.Saturday},
			Expected: Week{weekStart: StartDefault, weekEnd: EndDefault},
		},
		"not enough days invalid": {
			Input:    []time.Weekday{time.Sunday, time.Friday},
			Expected: Week{weekStart: StartDefault, weekEnd: EndDefault},
		},
		"too few days given invalid": {
			Input:    []time.Weekday{time.Sunday},
			Expected: Week{weekStart: StartDefault, weekEnd: EndDefault},
		},
		"too many days use only first 2": { // extra days are ignored
			Input:    []time.Weekday{time.Sunday, time.Saturday, time.Monday},
			Expected: Week{weekStart: time.Sunday, weekEnd: time.Saturday},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestLastFullWeek(t *testing.T) {
	d := NewWeek(time.Monday, time.Sunday)
	fn := func(in time.Time) (output, error) {
		start, end := d.LastFullWeek(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"normal date": {
			Input: Date(2024, 02, 05),
			Expected: output{
				start: Date(2024, 01, 29),
				end:   Date(2024, 02, 04),
			},
		},
		"previous year": {
			Input: Date(2024, 01, 03),
			Expected: output{
				start: Date(2023, 12, 25),
				end:   Date(2023, 12, 31),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestLastFullWeekSundayStart(t *testing.T) {
	d := NewWeek(time.Sunday, time.Saturday)
	fn := func(in time.Time) (output, error) {
		start, end := d.LastFullWeek(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"normal date": {
			Input: Date(2024, 02, 05),
			Expected: output{
				start: Date(2024, 01, 28),
				end:   Date(2024, 02, 03),
			},
		},
		"previous year": {
			Input: Date(2024, 01, 03),
			Expected: output{
				start: Date(2023, 12, 24),
				end:   Date(2023, 12, 30),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestDate(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		h := Date(in.Year(), in.Month(), in.Day())
		return h, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"date and time": {
			Input:    time.Date(2020, 11, 15, 12, 5, 5, 5, time.FixedZone("UTC-7", -6*56*34)),
			Expected: time.Date(2020, 11, 15, 0, 0, 0, 0, time.UTC),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestFirstDayOfMonth(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		h := StartOfMonth(in)
		return h, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"normal date": {
			Input:    time.Date(2020, 11, 15, 15, 5, 5, 5, time.UTC),
			Expected: Date(2020, 11, 1),
		},
		"weird date": {
			Input:    time.Date(2020, 2, 30, 15, 5, 5, 5, time.UTC),
			Expected: Date(2020, 3, 1),
		},
		"another weird date": {
			Input:    time.Date(2020, 0, 0, 15, 5, 5, 5, time.UTC),
			Expected: Date(2019, 11, 1),
		},
	}
	trial.New(fn, cases).SubTest(t)
}

func TestPrevMonthToDate(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := PrevMonthToDate(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"leap day": {
			Input: Date(2024, 3, 31),
			Expected: output{
				start: Date(2024, 2, 1),
				end:   Date(2024, 2, 29),
			},
		},
		"smaller prev month": {
			Input: Date(2024, 5, 31),
			Expected: output{
				start: Date(2024, 4, 1),
				end:   Date(2024, 4, 30),
			},
		},
		"prev year": {
			Input: Date(2024, 1, 15),
			Expected: output{
				start: Date(2023, 12, 1),
				end:   Date(2023, 12, 15),
			},
		},
		"normal date": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2024, 1, 1),
				end:   Date(2024, 1, 15),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestWeekAddStartOfWeek(t *testing.T) {
	d := NewWeek(time.Monday, time.Sunday)
	type input struct {
		date  time.Time
		weeks int
	}
	fn := func(in input) (time.Time, error) {
		h := WeekAdd(in.date, in.weeks)
		y := d.StartOfWeek(h)
		return y, nil
	}

	cases := trial.Cases[input, time.Time]{
		"normal day": {
			Input:    input{Date(2024, 6, 26), -25},
			Expected: Date(2024, 1, 1),
		},
		"previous year": {
			Input:    input{Date(2023, 1, 1), 0},
			Expected: Date(2022, 12, 26),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestPrevYearMtd(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := PrevYearMtd(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"leap day": {
			Input: Date(2024, 2, 29),
			Expected: output{
				start: Date(2023, 2, 1),
				end:   Date(2023, 2, 28),
			},
		},
		"normal date": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2023, 2, 1),
				end:   Date(2023, 2, 15),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestPreviousYearToDate(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := PreviousYearToDate(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"leap day": {
			Input: Date(2024, 2, 29),
			Expected: output{
				start: Date(2023, 1, 1),
				end:   Date(2023, 2, 28),
			},
		},
		"normal date": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2023, 1, 1),
				end:   Date(2023, 2, 15),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestWeekAdd(t *testing.T) {
	type input struct {
		date  time.Time
		weeks int
	}
	fn := func(in input) (time.Time, error) {
		h := WeekAdd(in.date, in.weeks)
		return h, nil
	}

	cases := trial.Cases[input, time.Time]{
		"one week forward": {
			Input:    input{Date(2024, 6, 26), 1},
			Expected: Date(2024, 7, 3),
		},
		"two weeks back": {
			Input:    input{Date(2024, 2, 29), -2},
			Expected: Date(2024, 2, 15),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestPriorLastFullWeek(t *testing.T) {
	d := NewWeek(time.Monday, time.Sunday)
	fn := func(in time.Time) (output, error) {
		start, end := d.PriorLastFullWeek(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"day one": {
			Input: Date(2024, 3, 20),
			Expected: output{
				start: Date(2024, 3, 4),
				end:   Date(2024, 3, 10),
			},
		},
		"day two": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2024, 1, 29),
				end:   Date(2024, 2, 4),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestPrevYearLastFullWeek(t *testing.T) {
	d := NewWeek(time.Monday, time.Sunday)
	fn := func(in time.Time) (output, error) {
		start, end := d.PrevYearLastFullWeek(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"day one": {
			Input:    Date(2024, 3, 31),
			Expected: output{start: Date(2023, 3, 20), end: Date(2023, 3, 26)},
		},
		"day two": {
			Input:    Date(2024, 2, 15),
			Expected: output{start: Date(2023, 2, 6), end: Date(2023, 2, 12)},
		},
		"day three": {
			Input:    Date(2023, 1, 18),
			Expected: output{start: Date(2022, 1, 10), end: Date(2022, 1, 16)},
		},
	}
	trial.New(fn, cases).SubTest(t)
}

func TestYearToDate(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := YearToDate(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"day one": {
			Input: Date(2024, 3, 31),
			Expected: output{
				start: Date(2024, 1, 1),
				end:   Date(2024, 3, 31),
			},
		},
		"day two": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2024, 1, 1),
				end:   Date(2024, 2, 15),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestFullMonth(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := FullMonth(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"day one": {
			Input: Date(2024, 3, 15),
			Expected: output{
				start: Date(2024, 3, 1),
				end:   Date(2024, 3, 31),
			},
		},
		"leap year": {
			Input: Date(2024, 2, 15),
			Expected: output{
				start: Date(2024, 2, 1),
				end:   Date(2024, 2, 29),
			},
		},
		"year end": {
			Input: Date(2024, 12, 15),
			Expected: output{
				start: Date(2024, 12, 1),
				end:   Date(2024, 12, 31),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestPrevMonth(t *testing.T) {
	fn := func(in time.Time) (output, error) {
		start, end := PrevMonth(in)
		return output{
			start: start,
			end:   end,
		}, nil
	}

	cases := trial.Cases[time.Time, output]{
		"day one": {
			Input: Date(2024, 3, 15),
			Expected: output{
				start: Date(2024, 2, 1),
				end:   Date(2024, 2, 29),
			},
		},
		"leap year": {
			Input: Date(2024, 4, 15),
			Expected: output{
				start: Date(2024, 3, 1),
				end:   Date(2024, 3, 31),
			},
		},
		"new year": {
			Input: Date(2024, 1, 15),
			Expected: output{
				start: Date(2023, 12, 1),
				end:   Date(2023, 12, 31),
			},
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestDay(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		h := Day(in)
		return h, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"normal date": {
			Input:    time.Date(2020, 11, 15, 15, 5, 5, 5, time.UTC),
			Expected: Date(2020, 11, 15),
		},
		"weird date": {
			Input:    time.Date(2020, 2, 30, 15, 5, 5, 5, time.UTC),
			Expected: Date(2020, 2, 30),
		},
		"another weird date": {
			Input:    time.Date(2020, 0, 0, 15, 5, 5, 5, time.UTC),
			Expected: Date(2019, 11, 30),
		},
	}
	trial.New(fn, cases).SubTest(t)
}

func TestLastDayOfMonth(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		h := LastDayOfMonth(in)
		return h, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"normal date": {
			Input:    time.Date(2024, 11, 15, 15, 5, 5, 5, time.UTC),
			Expected: Date(2024, 11, 30),
		},
		"weird date": {
			Input:    time.Date(2024, 2, 30, 15, 5, 5, 5, time.UTC),
			Expected: Date(2024, 3, 31),
		},
		"another weird date": {
			Input:    time.Date(2024, 0, 0, 15, 5, 5, 5, time.UTC),
			Expected: Date(2023, 11, 30),
		},
	}
	trial.New(fn, cases).SubTest(t)
}
