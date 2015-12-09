package dtf

import (
	"errors"
	"time"
)

// Parse generate time.Time from W3C-DTF string
func Parse(timeStr string) (time.Time, error) {
	switch true {
	case IsYear(timeStr):
		return ParseYear(timeStr)
	case IsYearAndMonth(timeStr):
		return ParseYearAndMonth(timeStr)
	case IsCompleteDate(timeStr):
		return ParseCompleteDate(timeStr)
	case IsCompleteDateWithMinutes(timeStr):
		return ParseCompleteDateWithMinutes(timeStr)
	case IsCompleteDateWithSeconds(timeStr):
		return ParseCompleteDateWithSeconds(timeStr)
	case IsCompleteDateWithFractionOfSecond(timeStr):
		return ParseCompleteDateWithFractionOfSecond(timeStr)
	default:
		var parsedTime time.Time
		return parsedTime, errors.New("provided string is not W3C-DTF format")
	}
}
