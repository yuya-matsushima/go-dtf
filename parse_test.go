package dtf

import (
	"testing"
	"time"
)

type parseTest struct {
	str      string
	expected time.Time
}

func TestParse(t *testing.T) {
	parseTests := []parseTest{
		{
			"2015",
			time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2015-01",
			time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2015-01-15",
			time.Date(2015, time.January, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"2015-01-15T18:30+00:00",
			time.Date(2015, time.January, 15, 18, 30, 0, 0, time.UTC),
		},
		{
			"2015-01-15T18:30Z",
			time.Date(2015, time.January, 15, 18, 30, 0, 0, time.UTC),
		},
		{
			"2015-01-15T18:30:20+00:00",
			time.Date(2015, time.January, 15, 18, 30, 20, 0, time.UTC),
		},
		{
			"2015-01-15T18:30:20Z",
			time.Date(2015, time.January, 15, 18, 30, 20, 0, time.UTC),
		},
		{
			"2015-01-15T18:30:20.123456789+00:00",
			time.Date(2015, time.January, 15, 18, 30, 20, 123456789, time.UTC),
		},
		{
			"2015-01-15T18:30:20.123456789Z",
			time.Date(2015, time.January, 15, 18, 30, 20, 123456789, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := Parse(test.str)
		if err != nil {
			t.Errorf(
				"#%d Parse() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d Parse() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseYear(t *testing.T) {
	parseTests := []parseTest{
		{
			"2000",
			time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2015",
			time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2020",
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseYear(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseYear() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseYear() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseYearAndMonth(t *testing.T) {
	parseTests := []parseTest{
		{
			"2015-12",
			time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"2020-01",
			time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseYearAndMonth(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseYearAndMonth() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseYearAndMonth() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseCompleteDate(t *testing.T) {
	parseTests := []parseTest{
		{
			"2015-12-19",
			time.Date(2015, time.December, 19, 0, 0, 0, 0, time.UTC),
		},
		{
			"2020-01-10",
			time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseCompleteDate(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseCompleteDate() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseCompleteDate() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseCompleteDateWithMinutes(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	parseTests := []parseTest{
		{
			"2015-12-19T18:30+09:00",
			time.Date(2015, time.December, 19, 18, 30, 0, 0, location),
		},
		{
			"2020-01-10T09:10+09:00",
			time.Date(2020, time.January, 10, 9, 10, 0, 0, location),
		},
		{
			"2015-12-19T18:30+00:00",
			time.Date(2015, time.December, 19, 18, 30, 0, 0, time.UTC),
		},
		{
			"2015-12-19T18:30Z",
			time.Date(2015, time.December, 19, 18, 30, 0, 0, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseCompleteDateWithMinutes(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseCompleteDateWithMinutes() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseCompleteDateWithMinutes() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseCompleteDateWithSeconds(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	parseTests := []parseTest{
		{
			"2015-12-19T18:30:22+09:00",
			time.Date(2015, time.December, 19, 18, 30, 22, 0, location),
		},
		{
			"2020-01-09T09:10:08+09:00",
			time.Date(2020, time.January, 9, 9, 10, 8, 0, location),
		},
		{
			"2015-12-19T18:30:22+00:00",
			time.Date(2015, time.December, 19, 18, 30, 22, 0, time.UTC),
		},
		{
			"2015-12-19T18:30:22Z",
			time.Date(2015, time.December, 19, 18, 30, 22, 0, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseCompleteDateWithSeconds(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseCompleteDateWithSeconds() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseCompleteDateWithSeconds() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}

func TestParseCompleteDateWithFractionOfSecond(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Tokyo")
	parseTests := []parseTest{
		{
			"2015-12-19T18:30:22.123456789+09:00",
			time.Date(2015, time.December, 19, 18, 30, 22, 123456789, location),
		},
		{
			"2020-01-09T10:08:10.222333444+09:00",
			time.Date(2020, time.January, 9, 10, 8, 10, 222333444, location),
		},
		{
			"2015-12-19T18:30:22.123456789+00:00",
			time.Date(2015, time.December, 19, 18, 30, 22, 123456789, time.UTC),
		},
		{
			"2015-12-19T18:30:22.123456789Z",
			time.Date(2015, time.December, 19, 18, 30, 22, 123456789, time.UTC),
		},
	}

	for i, test := range parseTests {
		ret, err := ParseCompleteDateWithFractionOfSecond(test.str)
		if err != nil {
			t.Errorf(
				"#%d ParseCompleteDateWithFractionOfSecond() should not return error:%s if given %s",
				i,
				err.Error(),
				test.str,
			)
		}

		if ret.UnixNano() != test.expected.UnixNano() {
			t.Errorf(
				"#%d ParseCompleteDateWithFractionOfSecond() should return %s if given %s. actual:%s",
				i,
				test.expected.String(),
				test.str,
				ret.String(),
			)
		}
	}
}
