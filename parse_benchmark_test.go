package dtf

import "testing"

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("2015-01-15T18:30:20.123456789+00:00")
	}
}

func BenchmarkParseYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseYear("2015")
	}
}

func BenchmarkParseYearAndMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseYearAndMonth("2015-12")
	}
}

func BenchmarkParseCompleteDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCompleteDate("2015-12-09")
	}
}

func BenchmarkParseCompleteDateWithMinutes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCompleteDateWithMinutes("2015-12-19T18:30+09:00")
	}
}

func BenchmarkParseCompleteDateWithSeconds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCompleteDateWithSeconds("2015-12-19T18:30:22+09:00")
	}
}

func BenchmarkParseCompleteDateWithFractionOfSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseCompleteDateWithFractionOfSecond("2015-12-19T18:30:22.123456789+09:00")
	}
}
