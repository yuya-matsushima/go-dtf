package dtf

import (
	"fmt"
)

func ExampleParse() {
	parsed, _ := Parse("2015-12-09T14:22:30+09:00")
	fmt.Println(parsed.String())
	// Output: 2015-12-09 14:22:30 +0900 JST
}
