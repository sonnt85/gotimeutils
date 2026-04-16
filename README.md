# gotimeutils

Time utility library for Go — flexible time parsing, period boundaries, stepped sleep, and duration formatting.

## Installation

```bash
go get github.com/sonnt85/gotimeutils
```

## Features

- Parse human-readable strings into `time.Time` using 25+ built-in formats
- Compute period boundaries: beginning/end of minute, hour, day, week, month, quarter, year
- Check if a time falls between two string-specified times
- Stepped sleep with jitter: gradually increase sleep interval up to a maximum
- Stepped ticker: same behavior but as a `time.Ticker`
- Random sleep helpers: sleep for a random duration within a range
- Duration formatting including day units (`1d12h30m`)
- Timestamp and date string helpers (UTC, local, LA timezone)

## Usage

```go
import "github.com/sonnt85/gotimeutils"

// Parse flexible date strings
t, _ := gotimeutils.Parse("2024-01-15")
t2, _ := gotimeutils.Parse("Jan 15, 2024 at 3:04 PM")

// Period boundaries
start := gotimeutils.BeginningOfMonth()
end   := gotimeutils.EndOfMonth()

// Check range
inRange := gotimeutils.Between("2024-01-01", "2024-12-31")

// Stepped sleep (increases each call, resets at max)
ss := gotimeutils.NewSleepStep(100*time.Millisecond, 1*time.Second, 30*time.Second)
ss.Sleep() // sleeps 1s + jitter
ss.Sleep() // sleeps longer, up to 30s then resets

// Random sleep
gotimeutils.SleepRandRange(1*time.Second, 5*time.Second)

// Format duration with days
s := gotimeutils.StringDuration(25*time.Hour + 30*time.Minute) // "1d1h30m0s"
```

## API

### Parsing
- `Parse(strs ...string) (time.Time, error)` — parse one or more date/time strings
- `MustParse(strs ...string) time.Time` — parse or panic
- `ParseInLocation(loc, strs...)` — parse in a specific timezone
- `With(t time.Time) *Now` — wrap a time for method chaining
- `NewNow(t time.Time) *Now` — alias for `With`

### Period Boundaries (package-level)
- `BeginningOf{Minute,Hour,Day,Week,Month,Quarter,Year}() time.Time`
- `EndOf{Minute,Hour,Day,Week,Month,Quarter,Year}() time.Time`
- `Monday/Sunday(strs ...string) time.Time` — Monday/Sunday of current or given week
- `Between(time1, time2 string) bool` — check if now is between two times
- `Quarter() uint` — current quarter (1–4)
- `NumDaysOfMonth() int` — number of days in current month

### Stepped Sleep
- `NewSleepStep(step, min, max Duration) *SleepStep` — create stepped sleeper
- `(*SleepStep).Sleep()` — sleep and advance the step
- `NewTickerStep(step, min, max Duration) *TickerStep` — stepped ticker
- `(*TickerStep).C() <-chan time.Time` — ticker channel
- `(*TickerStep).Update()` — advance the step interval

### Helpers
- `SleepRandRange(min, max Duration)` — sleep random duration in range
- `SleepRandMax(max Duration)` — sleep up to max
- `StringDuration(d Duration) string` — format duration with day units
- `ConvertTimestampsToLocalTime(unix int64) time.Time` — Unix timestamp to local time
- `TimeNowUTC() string` — current UTC time as `"2006-01-02 15:04:05"` string
- `GetTodaysDate/GetTodaysDateTime/GetTodaysDateTimeFormatted() string` — formatted current date/time

## License

MIT License - see [LICENSE](LICENSE) for details.
