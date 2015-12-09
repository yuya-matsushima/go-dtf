package dtf

import (
	"regexp"
)

var (
	year = "^[1-9][0-9]{3}$"

	yearAndMonth = "^[1-9][0-9]{3}-(1[0-2]|0[0-9])$"

	completeDate = "^[1-9][0-9]{3}-(1[0-2]|0[0-9])-([0-2][0-9]|3[0-1])$"

	withMinutes = "^[1-9][0-9]{3}-(1[0-2]|0[0-9])-([0-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):[0-5][0-9]([-+]([01][0-9]|2[0-4])00|[A-Z]{3})$"

	withSeconds = "^[1-9][0-9]{3}-(1[0-2]|0[0-9])-([0-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60)([-+]([01][0-9]|2[0-4])00|[A-Z]{3})$"

	withFractionOfSecond = "^[1-9][0-9]{3}-(1[0-2]|0[0-9])-([0-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):[0-5][0-9]:([0-5][0-9]|60).[0-9]+([-+]([01][0-9]|2[0-4])00|[A-Z]{3})$"
)

// IsYear check timeStr is 'YYYY'
func IsYear(timeStr string) bool {
	match, _ := regexp.MatchString(year, timeStr)
	return match
}

// IsYearAndMonth check timeStr is 'YYYY-MM'
func IsYearAndMonth(timeStr string) bool {
	match, _ := regexp.MatchString(yearAndMonth, timeStr)
	return match
}

// IsCompleteDate check timeStr is 'YYYY-MM-DD'
func IsCompleteDate(timeStr string) bool {
	match, _ := regexp.MatchString(completeDate, timeStr)
	return match
}

// IsCompleteDateWithMinutes check timeStr is 'YYYY-MM-DDThh:mmTZD'
func IsCompleteDateWithMinutes(timeStr string) bool {
	match, _ := regexp.MatchString(withMinutes, timeStr)
	return match
}

// IsCompleteDateWithSeconds check timeStr is 'YYYY-MM-DDThh:mm:ssTZD'
func IsCompleteDateWithSeconds(timeStr string) bool {
	match, _ := regexp.MatchString(withSeconds, timeStr)
	return match
}

// IsCompleteDateWithFractionOfSecond check timeStr is 'YYYY-MM-DDThh:mm:ss.sTZD'
func IsCompleteDateWithFractionOfSecond(timeStr string) bool {
	match, _ := regexp.MatchString(withFractionOfSecond, timeStr)
	return match
}
