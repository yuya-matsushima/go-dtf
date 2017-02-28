package dtf

import (
	"testing"
)

type evTest struct {
	str      string
	expected bool
}

func TestIsYear(t *testing.T) {
	evTests := []evTest{
		{"2015", true},
		{"2015-12", false},
	}

	for i, test := range evTests {
		if IsYear(test.str) != test.expected {
			t.Errorf(
				"%d IsYear() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsYearAndMonth(t *testing.T) {
	evTests := []evTest{
		{"2015-12", true},
		{"2015", false},
		{"2015-13", false},
		{"2015-20", false},
	}

	for i, test := range evTests {
		if IsYearAndMonth(test.str) != test.expected {
			t.Errorf(
				"%d IsYearAndMonth() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsCompleteDate(t *testing.T) {
	evTests := []evTest{
		{"2015-12-09", true},
		{"2015-15-01", false},
		{"2015-12", false},
		{"2015-12-32", false},
	}

	for i, test := range evTests {
		if IsCompleteDate(test.str) != test.expected {
			t.Errorf(
				"%d IsCompleteDate() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsCompleteDateWithMinutes(t *testing.T) {
	evTests := []evTest{
		{"2015-12-09T16:20+09:00", true},
		{"2015-12-09T16:20-09:00", true},
		{"2015-12-09T16:20Z", true},
		{"2015-12-09T16:20+25:00", false},
		{"2015-12-09T16:20A", false},
		{"2015-12-09T16:20z", false},
	}

	for i, test := range evTests {
		if IsCompleteDateWithMinutes(test.str) != test.expected {
			t.Errorf(
				"%d IsCompleteDateWithMinutes() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsCompleteDateWithSeconds(t *testing.T) {
	evTests := []evTest{
		{"2015-12-09T16:20:30+09:00", true},
		{"2015-12-09T16:20:30Z", true},
		{"2015-12-09T16:20:60+09:00", true},
		{"2015-12-09T16:20:61+09:00", false},
	}

	for i, test := range evTests {
		if IsCompleteDateWithSeconds(test.str) != test.expected {
			t.Errorf(
				"%d IsCompleteDateWithSeconds() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsCompleteDateWithFractionOfSecond(t *testing.T) {
	evTests := []evTest{
		{"2015-12-09T16:20:30.123456789+09:00", true},
		{"2015-12-09T16:20:30.45Z", true},
		{"2015-12-09T16:20:30.123456789+09:00", true},
		{"2015-12-09T16:20:30.9999999999+09:00", true},
		{"2015-12-09T16:20:600.9999999999+09:00", false},
	}

	for i, test := range evTests {
		if IsCompleteDateWithFractionOfSecond(test.str) != test.expected {
			t.Errorf(
				"%d IsCompleteDateWithFractionOfSecond() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsUTC(t *testing.T) {
	evTests := []evTest{
		{"Z", true},
		{"+09:00", false},
	}

	for i, test := range evTests {
		if IsUTC(test.str) != test.expected {
			t.Errorf(
				"%d IsUTC() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}

func TestIsW3CDTF(t *testing.T) {
	evTests := []evTest{
		{"2015-12-09T16:20:30.123456789+09:00", true},
		{"20151209142233", false},
	}

	for i, test := range evTests {
		if IsW3CDTF(test.str) != test.expected {
			t.Errorf(
				"%d IsW3CDTF() should return %t if given %s",
				i,
				test.expected,
				test.str,
			)
		}
	}
}
