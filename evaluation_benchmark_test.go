package dtf

import "testing"

func BenchmarkIsYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsYear("2015")
	}
}

func BenchmarkIsYearAndMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsYearAndMonth("2015-12")
	}
}

func BenchmarkIsCompleteDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsCompleteDate("2015-12-10")
	}
}

func BenchmarkIsCompleteDateWithMinutes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsCompleteDateWithMinutes("2015-12-09T16:20+09:00")
	}
}

func BenchmarkIsCompleteDateWithSeconds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsCompleteDateWithSeconds("2015-12-09T16:20:30+09:00")
	}
}

func BenchmarkIsIsCompleteDateWithFractionOfSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsCompleteDateWithFractionOfSecond("2015-12-09T16:20:30.123456789+09:00")
	}
}

func BenchmarkIsUTC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsUTC("Z")
	}
}

func BenchmarkIsW3CDTF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsW3CDTF("2015-12-09T16:20:30.123456789+09:00")
	}
}
