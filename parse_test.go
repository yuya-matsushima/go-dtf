package dtf

import (
	"testing"
	"time"
)

func TestParseYear(t *testing.T) {
	expected := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseYear("2015")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseYear return unexpected time object:" + result.String())
	}
}

func TestParseYearAndMonth(t *testing.T) {
	expected := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseYearAndMonth("2015-12")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseYearAndMonth return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDate(t *testing.T) {
	expected := time.Date(2015, time.December, 19, 0, 0, 0, 0, time.UTC)
	result, err := ParseCompleteDate("2015-12-19")

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Error("ParseCompleteDate return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithMinutes(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 0, 0, location)
	result, err := ParseCompleteDateWithMinutes("2015-12-19T18:30+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithMinutesContainsTimezoneString(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 0, 0, location)
	result, err := ParseCompleteDateWithMinutes("2015-12-19T18:30JST")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithSeconds(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 0, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithSecondsContainsTimezoneString(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 0, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22JST")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithMinutes return unexpected time object:" + result.String())
	}
}

func TestParseCompleteDateWithFractionOfSecond(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 123456789, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22.123456789+09:00")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithFractionOfSecond return unexpected time object:" + result.String())
		t.Error(result)
		t.Error(expected)
	}
}

func TestParseCompleteDateWithFractionOfSecondContainsTimezoneString(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2015, time.December, 19, 18, 30, 22, 123456789, location)
	result, err := ParseCompleteDateWithSeconds("2015-12-19T18:30:22.123456789JST")

	if err != nil {
		t.Error(err)
	}

	if result.String() != expected.String() {
		t.Error("ParseCompleteDateWithFractionOfSecond return unexpected time object:" + result.String())
		t.Error(result)
		t.Error(expected)
	}
}
