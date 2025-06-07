package lotime

import (
	"fmt"
	"time"
)

// GetWeekOfMonth returns the week number as int in month by date
func GetWeekOfMonth(dt time.Time, isFirstDayMon bool) int {
	// Get the first day of the month
	firstDayOfMonth := time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, dt.Location())

	// Get offset for first weekday
	adjustedWeekday := 0
	if isFirstDayMon {
		adjustedWeekday = (int(firstDayOfMonth.Weekday()) + 6) % 7
	}

	// Day of month
	dayOfMonth := dt.Day()

	// Calculate adjusted day
	adjustedDay := dayOfMonth + adjustedWeekday - 1

	// Calculate week number
	return (adjustedDay / 7) + 1
}

// GetDayOfYear return day of year by date
func GetDayOfYear(year int, month int, day int) int {
	monthName := time.Month(month)
	date := time.Date(year, monthName, day, 0, 0, 0, 0, time.UTC)
	return date.YearDay()
}

// DateIsPassedInCurrentYear Проверяем, прошла ли дата в текущем году
func DateIsPassedInCurrentYear(dt time.Time) bool {
	now := time.Now().Local()

	date := time.Date(now.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.Local)

	if now.After(date) {
		return true
	}
	return false
}

// DateYearsSince Amount of years since the given date
func DateYearsSince(dt time.Time) int {
	now := time.Now()

	years := now.Year() - dt.Year()

	// If current date is before dt by day or month, then year is not complete
	if now.Month() < dt.Month() || (now.Month() == dt.Month() && now.Day() < dt.Day()) {
		years--
	}

	return years
}

// NthOrLastWeekdayOfMonth Calculate the nth or last weekday of a given month and year.
func NthOrLastWeekdayOfMonth(year int, month time.Month, userWeekday int, weekNum int, isFirstDayMon bool) (time.Time, error) {
	if userWeekday < 1 || userWeekday > 7 {
		return time.Time{}, fmt.Errorf("invalid weekday: %d", userWeekday)
	}
	if weekNum < 0 || weekNum > 5 {
		return time.Time{}, fmt.Errorf("invalid week number: %d", weekNum)
	}

	targetWeekday := time.Weekday(0)
	if isFirstDayMon {
		targetWeekday = time.Weekday(userWeekday % 7) // Sunday=0 ... Saturday=6
	}

	if weekNum == 0 {
		// Search for the last weekday of the month
		firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
		lastDayOfMonth := firstOfNextMonth.AddDate(0, 0, -1)

		for lastDayOfMonth.Weekday() != targetWeekday {
			lastDayOfMonth = lastDayOfMonth.AddDate(0, 0, -1)
		}
		return lastDayOfMonth, nil
	}

	// Search for the nth weekday of the month
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	offset := (int(targetWeekday) - int(firstOfMonth.Weekday()) + 7) % 7
	day := 1 + offset + (weekNum-1)*7

	// Check if the date is valid
	if day > 31 || time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Month() != month {
		return time.Time{}, fmt.Errorf("no such weekday occurrence in this month")
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}
