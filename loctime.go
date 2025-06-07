package lotime

import (
	"sync"
	"time"
)

var (
	globalLocation *time.Location
	mu             sync.RWMutex
	defaultLoc     = time.UTC
)

func init() {
	globalLocation = defaultLoc
}

// SetGlobalTimeZone sets the default timezone
func SetGlobalTimeZone(name string) error {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()
	globalLocation = loc
	return nil
}

func Date(year int, month time.Month, day, hour, minute, second, nSecond int) time.Time {
	return time.Date(year, month, day, hour, minute, second, nSecond, globalLocation)
}

// SetFixedTimeZone set timezone to fixed value with specified hours offset
func SetFixedTimeZone(name string, hours int) {
	loc := time.FixedZone(name, hours*60*60)
	mu.Lock()
	defer mu.Unlock()
	globalLocation = loc
}

// Reset resets timezone to UTC.
func Reset() {
	mu.Lock()
	defer mu.Unlock()
	globalLocation = defaultLoc
}

// Now return time in current timezone
func Now() time.Time {
	mu.RLock()
	defer mu.RUnlock()
	return time.Now().In(globalLocation)
}

// Parse date string in global timezone
//
//	Year: "2006" "06"
//	Month: "Jan" "January" "01" "1"
//	Day of the week: "Mon" "Monday"
//	Day of the month: "2" "_2" "02"
//	Day of the year: "__2" "002"
//	Hour: "15" "3" "03" (PM or AM)
//	Minute: "4" "04"
//	Second: "5" "05"
//	AM/PM mark: "PM"
//
// Numeric time zone offsets format as follows:
//
//	"-0700"     ±hhmm
//	"-07:00"    ±hh:mm
//	"-07"       ±hh
//	"-070000"   ±hhmmss
//	"-07:00:00" ±hh:mm:ss
//
// Replacing the sign in the format with a Z triggers
// the ISO 8601 behavior of printing Z instead of an
// offset for the UTC zone. Thus:
//
//	"Z0700"      Z or ±hhmm
//	"Z07:00"     Z or ±hh:mm
//	"Z07"        Z or ±hh
//	"Z070000"    Z or ±hhmmss
//	"Z07:00:00"  Z or ±hh:mm:ss
func Parse(layout, value string) (time.Time, error) {
	mu.RLock()
	defer mu.RUnlock()
	return time.ParseInLocation(layout, value, globalLocation)
}

func Unix() int64 {
	return int64(Now().UnixNano() / 1e9)
}

// Location return current timezone
func Location() *time.Location {
	mu.RLock()
	defer mu.RUnlock()
	return globalLocation
}
