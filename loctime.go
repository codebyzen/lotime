package github.com/codebyzen/lotime

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

// Parse parse date string in global timezone
func Parse(layout, value string) (time.Time, error) {
	mu.RLock()
	defer mu.RUnlock()
	return time.ParseInLocation(layout, value, globalLocation)
}

// Location return current timezone
func Location() *time.Location {
	mu.RLock()
	defer mu.RUnlock()
	return globalLocation
}
