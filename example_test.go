package dtf

import (
	"fmt"
)

func ExampleParse() {
	parsed, _ := Parse("2015-12-09T14:22:30+09:00")
	fmt.Println(parsed.Format("2006-01-02T15:04:05-07:00"))
	// Output: 2015-12-09T14:22:30+09:00
}
