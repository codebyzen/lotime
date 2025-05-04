# lotime
Easily manage global time zones across your Go application. 
Set a default timezone (like `America/New_York`) once, and all time operations (`Now()`, `Parse()`, etc.) with lotime will automatically use it. 

No more repetitive `In(location)` calls!

```go
import "github.com/codebyzen/lotime"

// Set global timezone (e.g., New-York)
lotime.SetGlobalTimeZone("America/New_York") 

// All operations now use Boston time:
now := lotime.Now()                                                     // Current time in New-York
parsed, _ := lotime.Parse("2006-01-02T15:04:05", "2025-05-05T10:00:00") // Parsed in New-York
```

### ✨ Features
- Set a global timezone (e.g., "America/New_York", "UTC-4") for your entire app.
- Thread-safe operations (uses sync.RWMutex under the hood).
- Zero-dependency (only relies on Go’s standard time package).
- Batteries included:
  - Now() – Current time in global timezone.
  - Parse(layout, value) – Parse time in global timezone.
  - SetFixedTimeZone() – Support for custom offsets (e.g., UTC+5:30).
  - Location() – Get the current global timezone.

### 🚀 Use Cases
- Apps where all timestamps must use a specific timezone.
- Simplify code by avoiding repetitive In(location) calls.
- Testing (override timezone globally for consistent results).

### 📌 Why lotime?
- No magic: Explicitly set a timezone instead of relying on time.Local.
- Lightweight: ~70 LOC, no external dependencies.
- Idiomatic: Wraps Go’s time package seamlessly.

### 🔧 Installation
```bash
go get github.com/codebyzen/lotime
```

### 📄 Example
```go
package main

import (
	"fmt"
	"github.com/codebyzen/lotime"
)

func main() {
	// Set global timezone to Tokyo
	lotime.SetGlobalTimeZone("Asia/Tokyo")

	// All operations now use Tokyo time
	fmt.Println("Current time in Tokyo:", lotime.Now())

	// Parse a date in Tokyo time
	tokyoTime, _ := lotime.Parse("2006-01-02", "2024-12-31")
	fmt.Println("Parsed time:", tokyoTime)
}
```


### ⚠️ Caveats
- Not recommended for apps needing per-request timezones (use context instead).
- Timezone names must be valid IANA identifiers (e.g., "Europe/London"). 
  - But you can use `lotime.SetFixedTimeZone("Custom", 5)` to set timezone by hours offset instead of naming.

Contributions welcome! 🛠️

This project is licensed under the MIT License - see the LICENSE file for details.
