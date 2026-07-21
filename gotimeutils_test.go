package gotimeutils

import (
	"testing"
	"time"
)

// fixedTime returns a deterministic time for testing: 2024-03-15 14:30:45.123456789 UTC (Friday)
func fixedTime() time.Time {
	return time.Date(2024, time.March, 15, 14, 30, 45, 123456789, time.UTC)
}

func TestConvertTimestampsToLocalTime(t *testing.T) {
	// Unix epoch 0 should be 1970-01-01 00:00:00 UTC (in local timezone)
	ts := int64(0)
	result := ConvertTimestampsToLocalTime(ts)
	// Verify the underlying Unix timestamp is preserved
	if result.Unix() != 0 {
		t.Errorf("expected Unix() == 0, got %d", result.Unix())
	}

	// A known timestamp: 1710510645 = 2024-03-15 13:50:45 UTC
	ts2 := int64(1710510645)
	result2 := ConvertTimestampsToLocalTime(ts2)
	if result2.Unix() != ts2 {
		t.Errorf("expected Unix() == %d, got %d", ts2, result2.Unix())
	}
	utcResult := result2.UTC()
	if utcResult.Year() != 2024 || utcResult.Month() != time.March || utcResult.Day() != 15 {
		t.Errorf("expected 2024-03-15 UTC, got %v", utcResult)
	}
	if utcResult.Hour() != 13 || utcResult.Minute() != 50 || utcResult.Second() != 45 {
		t.Errorf("expected 13:50:45 UTC, got %v", utcResult)
	}

	// Negative timestamp (before epoch)
	tsNeg := int64(-86400)
	resultNeg := ConvertTimestampsToLocalTime(tsNeg)
	if resultNeg.Unix() != tsNeg {
		t.Errorf("expected Unix() == %d, got %d", tsNeg, resultNeg.Unix())
	}
}

func TestBeginningOfMinute(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfMinute()

	expected := time.Date(2024, time.March, 15, 14, 30, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfMinute: expected %v, got %v", expected, result)
	}
}

func TestBeginningOfHour(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfHour()

	expected := time.Date(2024, time.March, 15, 14, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfHour: expected %v, got %v", expected, result)
	}
}

func TestBeginningOfDay(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfDay()

	expected := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfDay: expected %v, got %v", expected, result)
	}
}

func TestBeginningOfWeek(t *testing.T) {
	// fixedTime is Friday 2024-03-15, default WeekStartDay is Sunday
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfWeek()

	// Sunday of that week is 2024-03-10
	expected := time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfWeek (Sunday start): expected %v, got %v", expected, result)
	}

	// Test with Monday as week start
	config := &Config{
		WeekStartDay: time.Monday,
		TimeFormats:  TimeFormats,
	}
	nowMon := config.With(ft)
	resultMon := nowMon.BeginningOfWeek()
	// Monday of that week is 2024-03-11
	expectedMon := time.Date(2024, time.March, 11, 0, 0, 0, 0, time.UTC)
	if !resultMon.Equal(expectedMon) {
		t.Errorf("BeginningOfWeek (Monday start): expected %v, got %v", expectedMon, resultMon)
	}
}

func TestBeginningOfMonth(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfMonth()

	expected := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfMonth: expected %v, got %v", expected, result)
	}
}

func TestBeginningOfYear(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.BeginningOfYear()

	expected := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("BeginningOfYear: expected %v, got %v", expected, result)
	}
}

func TestEndOfMinute(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfMinute()

	expected := time.Date(2024, time.March, 15, 14, 30, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfMinute: expected %v, got %v", expected, result)
	}
}

func TestEndOfHour(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfHour()

	expected := time.Date(2024, time.March, 15, 14, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfHour: expected %v, got %v", expected, result)
	}
}

func TestEndOfDay(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfDay()

	expected := time.Date(2024, time.March, 15, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfDay: expected %v, got %v", expected, result)
	}
}

func TestEndOfMonth(t *testing.T) {
	// March has 31 days
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfMonth()

	expected := time.Date(2024, time.March, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfMonth (March): expected %v, got %v", expected, result)
	}

	// February in leap year 2024 has 29 days
	febTime := time.Date(2024, time.February, 10, 12, 0, 0, 0, time.UTC)
	nowFeb := With(febTime)
	resultFeb := nowFeb.EndOfMonth()
	expectedFeb := time.Date(2024, time.February, 29, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !resultFeb.Equal(expectedFeb) {
		t.Errorf("EndOfMonth (Feb leap year): expected %v, got %v", expectedFeb, resultFeb)
	}

	// February in non-leap year 2023 has 28 days
	febTimeNonLeap := time.Date(2023, time.February, 10, 12, 0, 0, 0, time.UTC)
	nowFebNL := With(febTimeNonLeap)
	resultFebNL := nowFebNL.EndOfMonth()
	expectedFebNL := time.Date(2023, time.February, 28, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !resultFebNL.Equal(expectedFebNL) {
		t.Errorf("EndOfMonth (Feb non-leap): expected %v, got %v", expectedFebNL, resultFebNL)
	}
}

func TestEndOfYear(t *testing.T) {
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfYear()

	expected := time.Date(2024, time.December, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfYear: expected %v, got %v", expected, result)
	}
}

func TestParse(t *testing.T) {
	// Use a fixed base time
	base := time.Date(2024, time.March, 15, 14, 30, 45, 0, time.UTC)
	now := With(base)

	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "full date",
			input:    "2024-3-15",
			expected: time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "date with time",
			input:    "2024-3-15 10:20:30",
			expected: time.Date(2024, time.March, 15, 10, 20, 30, 0, time.UTC),
		},
		{
			name:     "year only",
			input:    "2023",
			expected: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "year-month",
			input:    "2024-1",
			expected: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "date with hour only",
			input:    "2024-3-15 10",
			expected: time.Date(2024, time.March, 15, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := now.Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", tt.input, err)
			}
			if !result.Equal(tt.expected) {
				t.Errorf("Parse(%q): expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}

	// Test parse error for invalid string
	_, err := now.Parse("not-a-date")
	if err == nil {
		t.Error("expected error for invalid date string, got nil")
	}
}

func TestBetween(t *testing.T) {
	// Time is 14:30:45, check if between 14:00 and 15:00
	base := time.Date(2024, time.March, 15, 14, 30, 45, 0, time.UTC)
	now := With(base)

	if !now.Between("14:00", "15:00") {
		t.Error("expected 14:30:45 to be between 14:00 and 15:00")
	}

	if now.Between("15:00", "16:00") {
		t.Error("expected 14:30:45 NOT to be between 15:00 and 16:00")
	}

	if now.Between("13:00", "14:00") {
		t.Error("expected 14:30:45 NOT to be between 13:00 and 14:00")
	}

	// Test with full date-time strings
	if !now.Between("2024-3-15 14:00", "2024-3-15 15:00") {
		t.Error("expected time to be between 2024-3-15 14:00 and 2024-3-15 15:00")
	}

	// Edge case: exactly at boundary should return false (strict comparison)
	exactBase := time.Date(2024, time.March, 15, 14, 0, 0, 0, time.UTC)
	nowExact := With(exactBase)
	if nowExact.Between("14:00", "15:00") {
		t.Error("expected exact boundary time NOT to be between (strict inequality)")
	}
}

func TestQuarter(t *testing.T) {
	tests := []struct {
		month    time.Month
		expected uint
	}{
		{time.January, 1},
		{time.February, 1},
		{time.March, 1},
		{time.April, 2},
		{time.May, 2},
		{time.June, 2},
		{time.July, 3},
		{time.August, 3},
		{time.September, 3},
		{time.October, 4},
		{time.November, 4},
		{time.December, 4},
	}

	for _, tt := range tests {
		t.Run(tt.month.String(), func(t *testing.T) {
			tm := time.Date(2024, tt.month, 15, 12, 0, 0, 0, time.UTC)
			now := With(tm)
			result := now.Quarter()
			if result != tt.expected {
				t.Errorf("Quarter for %s: expected %d, got %d", tt.month, tt.expected, result)
			}
		})
	}
}

func TestMonday(t *testing.T) {
	// fixedTime is Friday 2024-03-15
	ft := fixedTime()
	now := With(ft)
	result := now.Monday()

	// Monday of that week is 2024-03-11
	expected := time.Date(2024, time.March, 11, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("Monday: expected %v, got %v", expected, result)
	}

	// Test when the day is Sunday (2024-03-10)
	sundayTime := time.Date(2024, time.March, 10, 10, 0, 0, 0, time.UTC)
	nowSun := With(sundayTime)
	resultSun := nowSun.Monday()
	// Monday of the previous week: 2024-03-04
	expectedSun := time.Date(2024, time.March, 4, 0, 0, 0, 0, time.UTC)
	if !resultSun.Equal(expectedSun) {
		t.Errorf("Monday (from Sunday): expected %v, got %v", expectedSun, resultSun)
	}

	// Test when the day is Monday itself (2024-03-11)
	mondayTime := time.Date(2024, time.March, 11, 10, 0, 0, 0, time.UTC)
	nowMon := With(mondayTime)
	resultMon := nowMon.Monday()
	expectedMon := time.Date(2024, time.March, 11, 0, 0, 0, 0, time.UTC)
	if !resultMon.Equal(expectedMon) {
		t.Errorf("Monday (from Monday): expected %v, got %v", expectedMon, resultMon)
	}
}

func TestSunday(t *testing.T) {
	// fixedTime is Friday 2024-03-15
	ft := fixedTime()
	now := With(ft)
	result := now.Sunday()

	// Sunday of that week is 2024-03-17
	expected := time.Date(2024, time.March, 17, 0, 0, 0, 0, time.UTC)
	if !result.Equal(expected) {
		t.Errorf("Sunday: expected %v, got %v", expected, result)
	}

	// Test when the day is already Sunday (2024-03-17)
	sundayTime := time.Date(2024, time.March, 17, 10, 0, 0, 0, time.UTC)
	nowSun := With(sundayTime)
	resultSun := nowSun.Sunday()
	// Sunday treats weekday=0 as 7, so (7-7)=0 days added => same day
	expectedSun := time.Date(2024, time.March, 17, 0, 0, 0, 0, time.UTC)
	if !resultSun.Equal(expectedSun) {
		t.Errorf("Sunday (from Sunday): expected %v, got %v", expectedSun, resultSun)
	}

	// Test when the day is Monday (2024-03-11)
	mondayTime := time.Date(2024, time.March, 11, 10, 0, 0, 0, time.UTC)
	nowMon := With(mondayTime)
	resultMon := nowMon.Sunday()
	expectedMon := time.Date(2024, time.March, 17, 0, 0, 0, 0, time.UTC)
	if !resultMon.Equal(expectedMon) {
		t.Errorf("Sunday (from Monday): expected %v, got %v", expectedMon, resultMon)
	}
}

func TestStringDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Duration
		expected string
	}{
		{
			name:     "less than a day",
			input:    2*time.Hour + 30*time.Minute + 15*time.Second,
			expected: "2h30m15s",
		},
		{
			name:     "exact days",
			input:    3 * 24 * time.Hour,
			expected: "3d",
		},
		{
			name:     "days and remainder",
			input:    2*24*time.Hour + 5*time.Hour + 30*time.Minute,
			expected: "2d5h30m0s",
		},
		{
			name:     "one day",
			input:    24 * time.Hour,
			expected: "1d",
		},
		{
			name:     "zero",
			input:    0,
			expected: "0s",
		},
		{
			name:     "negative less than a day",
			input:    -2 * time.Hour,
			expected: "2h0m0s", // sign is dropped for sub-day durations
		},
		{
			name:     "negative days",
			input:    -(3*24*time.Hour + 2*time.Hour),
			expected: "-3d2h0m0s",
		},
		{
			name:     "negative exact days",
			input:    -2 * 24 * time.Hour,
			expected: "-2d",
		},
		{
			name:     "milliseconds",
			input:    500 * time.Millisecond,
			expected: "500ms",
		},
		{
			name:     "one day plus one second",
			input:    24*time.Hour + time.Second,
			expected: "1d1s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringDuration(tt.input)
			if result != tt.expected {
				t.Errorf("StringDuration(%v): expected %q, got %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestNewNow(t *testing.T) {
	ft := fixedTime()
	now := NewNow(ft)
	if !now.Time.Equal(ft) {
		t.Errorf("NewNow: expected time %v, got %v", ft, now.Time)
	}
	// NewNow and With should produce equivalent results
	nowWith := With(ft)
	if !now.Time.Equal(nowWith.Time) {
		t.Errorf("NewNow and With should produce same time")
	}
}

func TestBeginningOfQuarter(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 (March)",
			input:    time.Date(2024, time.March, 15, 14, 30, 0, 0, time.UTC),
			expected: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 (May)",
			input:    time.Date(2024, time.May, 20, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 (September)",
			input:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 (December)",
			input:    time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := With(tt.input)
			result := now.BeginningOfQuarter()
			if !result.Equal(tt.expected) {
				t.Errorf("BeginningOfQuarter: expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEndOfQuarter(t *testing.T) {
	// Q1 ends March 31
	q1Time := time.Date(2024, time.February, 15, 12, 0, 0, 0, time.UTC)
	now := With(q1Time)
	result := now.EndOfQuarter()
	expected := time.Date(2024, time.March, 31, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfQuarter Q1: expected %v, got %v", expected, result)
	}

	// Q2 ends June 30
	q2Time := time.Date(2024, time.May, 10, 12, 0, 0, 0, time.UTC)
	nowQ2 := With(q2Time)
	resultQ2 := nowQ2.EndOfQuarter()
	expectedQ2 := time.Date(2024, time.June, 30, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !resultQ2.Equal(expectedQ2) {
		t.Errorf("EndOfQuarter Q2: expected %v, got %v", expectedQ2, resultQ2)
	}
}

func TestEndOfWeek(t *testing.T) {
	// fixedTime is Friday 2024-03-15, default WeekStartDay is Sunday
	ft := fixedTime()
	now := With(ft)
	result := now.EndOfWeek()

	// Week starts Sunday 2024-03-10, ends Saturday 2024-03-16 23:59:59.999999999
	expected := time.Date(2024, time.March, 16, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	if !result.Equal(expected) {
		t.Errorf("EndOfWeek: expected %v, got %v", expected, result)
	}
}
