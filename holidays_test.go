package dates

import (
	"testing"
	"time"

	"github.com/hydronica/trial"
)

func TestMartinLutherKingJrDay(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		output := MartinLutherKingJrDay(in)
		return output, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"2024": {
			Input:    Date(2024, 3, 20),
			Expected: Date(2024, 1, 15),
		},
		"2023": {
			Input:    Date(2023, 12, 20),
			Expected: Date(2023, 1, 16),
		},
		"2025": {
			Input:    Date(2025, 6, 20),
			Expected: Date(2025, 1, 20),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestNewYearsDay(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		output := NewYearsDay(in)
		return output, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"2024": {
			Input:    Date(2024, 3, 20),
			Expected: Date(2024, 1, 1),
		},
		"2023": {
			Input:    Date(2023, 12, 20),
			Expected: Date(2023, 1, 1),
		},
		"2025": {
			Input:    Date(2025, 6, 20),
			Expected: Date(2025, 1, 1),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestNewYearsEve(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		output := NewYearsEve(in)
		return output, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"2024": {
			Input:    Date(2024, 3, 20),
			Expected: Date(2024, 12, 31),
		},
		"2023": {
			Input:    Date(2023, 12, 20),
			Expected: Date(2023, 12, 31),
		},
		"2025": {
			Input:    Date(2025, 6, 20),
			Expected: Date(2025, 12, 31),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestMemorialDay(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		output := MemorialDay(in)
		return output, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"2024": {
			Input:    Date(2024, 3, 20),
			Expected: Date(2024, 5, 27),
		},
		"2023": {
			Input:    Date(2023, 12, 20),
			Expected: Date(2023, 5, 29),
		},
		"2025": {
			Input:    Date(2025, 6, 20),
			Expected: Date(2025, 5, 26),
		},
	}

	trial.New(fn, cases).SubTest(t)
}

func TestLaborDay(t *testing.T) {
	fn := func(in time.Time) (time.Time, error) {
		output := LaborDay(in)
		return output, nil
	}

	cases := trial.Cases[time.Time, time.Time]{
		"2024": {
			Input:    Date(2024, 3, 20),
			Expected: Date(2024, 9, 2),
		},
		"2023": {
			Input:    Date(2023, 12, 20),
			Expected: Date(2023, 9, 4),
		},
		"2025": {
			Input:    Date(2025, 6, 20),
			Expected: Date(2025, 9, 1),
		},
	}

	trial.New(fn, cases).SubTest(t)
}
