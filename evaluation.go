package dtf

import (
	"regexp"
	"strings"
)

const (
	timezone             = "([-+]([01][0-9]|2[0-4]):00|Z)"
	year                 = "[1-9][0-9]{3}"
	yearAndMonth         = year + "-(1[0-2]|0[0-9])"
	completeDate         = yearAndMonth + "-([0-2][0-9]|3[0-1])"
	hourMinutes          = "([01][0-9]|2[0-3]):[0-5][0-9]"
	withMinutes          = completeDate + "T" + hourMinutes + timezone
	withSeconds          = completeDate + "T" + hourMinutes + ":([0-5][0-9]|60)" + timezone
	withFractionOfSecond = completeDate + "T" + hourMinutes + ":([0-5][0-9]|60).[0-9]+" + timezone
)

var (
	rYear                 = regexp.MustCompile("^" + year + "$")
	rYearAndMonth         = regexp.MustCompile("^" + yearAndMonth + "$")
	rCompleteDate         = regexp.MustCompile("^" + completeDate + "$")
	rWithMinutes          = regexp.MustCompile("^" + withMinutes + "$")
	rWithSeconds          = regexp.MustCompile("^" + withSeconds + "$")
	rWithFractionOfSecond = regexp.MustCompile("^" + withFractionOfSecond + "$")
	rUTC                  = regexp.MustCompile("Z$")
)

// IsYear check timeStr is 'YYYY'
func IsYear(timeStr string) bool {
	return rYear.MatchString(timeStr)
}

// IsYearAndMonth check timeStr is 'YYYY-MM'
func IsYearAndMonth(timeStr string) bool {
	return rYearAndMonth.MatchString(timeStr)
}

// IsCompleteDate check timeStr is 'YYYY-MM-DD'
func IsCompleteDate(timeStr string) bool {
	return rCompleteDate.MatchString(timeStr)
}

// IsCompleteDateWithMinutes check timeStr is 'YYYY-MM-DDThh:mmTZD'
func IsCompleteDateWithMinutes(timeStr string) bool {
	return rWithMinutes.MatchString(timeStr)
}

// IsCompleteDateWithSeconds check timeStr is 'YYYY-MM-DDThh:mm:ssTZD'
func IsCompleteDateWithSeconds(timeStr string) bool {
	return rWithSeconds.MatchString(timeStr)
}

// IsCompleteDateWithFractionOfSecond check timeStr is 'YYYY-MM-DDThh:mm:ss.sTZD'
func IsCompleteDateWithFractionOfSecond(timeStr string) bool {
	return rWithFractionOfSecond.MatchString(timeStr)
}

// IsUTC check timeStr is UTC or not
func IsUTC(timeStr string) bool {
	return rUTC.MatchString(timeStr)
}

// IsW3CDTF check timeStr is match W3C-DTF format
func IsW3CDTF(timeStr string) bool {
	if !strings.Contains(timeStr, "T") {
		switch true {
		case IsYear(timeStr):
			return true
		case IsYearAndMonth(timeStr):
			return true
		case IsCompleteDate(timeStr):
			return true
		default:
			return false
		}
	} else {
		switch true {
		case IsCompleteDateWithMinutes(timeStr):
			return true
		case IsCompleteDateWithSeconds(timeStr):
			return true
		case IsCompleteDateWithFractionOfSecond(timeStr):
			return true
		default:
			return false
		}
	}
}
