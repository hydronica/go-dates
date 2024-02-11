# Dates Go Library
This is a simple Go library that provides helper functions and methods for calculating various date ranges and specific dates. It's designed to simplify date calculations and manipulations in your Go applications.

## Features

- **Week Struct**: Defines a week with a start and end day. You can create a new week with custom start and end days.
- **Date Function**: Creates a new date with the time truncated.
- **Day Function**: Returns the truncated date of the given time.
- **LastDayOfMonth Function**: Returns the last day of the month for a given date.
- **WeekAdd Function**: Adds or subtracts weeks from a given date.
- **StartOfWeek Method**: Returns the start of the week for a given date.
- **LastFullWeek Method**: Returns the start and end dates of the last full week.
- **PriorLastFullWeek Method**: Returns the start and end dates of the week prior to the last full week.
- **PrevYearLastFullWeek Method**: Returns the start and end dates of the last full week of the previous year.
- **MonthToDate Function**: Returns the 1st of the month to the given date.
- **FullMonth Function**: Returns the start and last day of the given date's month.
- **FirstOfNextMonth Function**: Returns the first day of the next month from a given date.
- **PrevMonth Function**: Returns the start and end dates of the previous month.
- **PrevMonthToDate Function**: Returns the start and end dates of the previous month to the given date.
- **PrevYearMtd Function**: Returns the start and end dates up to a given date of the same month in the previous year.
- **YearToDate Function**: Returns the start of the year and end date from a given date.
- **PreviousYearToDate Function**: Returns the start and end dates of the previous year for the given date.
- **StartOfMonth Function**: Returns the first day of the given date's month.

## todo:
- currently in the process of adding calulated holiday functions

## Usage

To use this library, import it in your Go application:

```go
import "dates"
```

Then, you can call any of the functions or methods provided by the library. For example, to get the start and end dates of the last full week:

```go
week := dates.NewWeek(time.Monday, time.Sunday)
start, end := week.LastFullWeek(time.Now())
```

## Contributing

Contributions are welcome! Please submit a pull request or create an issue to add new features or fix bugs.

## License

This library is licensed under the MIT License.
