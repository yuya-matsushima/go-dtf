package dtf

import (
	"time"
)

// ParseYear generate time.Time from 'YYYY' string
func ParseYear(timeStr string) (time.Time, error) {
	return time.Parse("2006", timeStr)
}

func ParseYearAndMonth(timeStr string) (time.Time, error) {
	return time.Parse("2006-01", timeStr)
}

func ParseCompleteDate(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02", timeStr)
}

func ParseCompleteDateWithMinutes(timeStr string) (time.Time, error) {
	if IsTimezoneString(timeStr) {
		return time.Parse("2006-01-02T15:04MST", timeStr)
	} else {
		return time.Parse("2006-01-02T15:04-07:00", timeStr)
	}
}

func ParseCompleteDateWithSeconds(timeStr string) (time.Time, error) {
	if IsTimezoneString(timeStr) {
		return time.Parse("2006-01-02T15:04:05MST", timeStr)
	} else {
		return time.Parse("2006-01-02T15:04:05-07:00", timeStr)
	}
}

func ParseCompleteDateWithFractionOfSecond(timeStr string) (time.Time, error) {
	if IsTimezoneString(timeStr) {
		return time.Parse("2006-01-02T15:04:05MST", timeStr)
	} else {
		return time.Parse("2006-01-02T15:04:05-07:00", timeStr)
	}
}
