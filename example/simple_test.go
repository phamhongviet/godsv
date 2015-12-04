package example

import (
	"fmt"

	"github.com/phamhongviet/godsv"
)

func ExampleMarshal() {
	row := godsv.Row{
		"this",
		"is",
		"a:simple",
		"test",
		"te\\st",
	}
	fmt.Println(godsv.Marshal(row))
	// Output:
	// this:is:a\:simple:test:te\\st
}

func ExampleUnmarshal() {
	fmt.Println(godsv.Unmarshal("this:is:a\\:simple:test:te\\\\st"))
	// Output:
	// [this is a:simple test te\st]
}
