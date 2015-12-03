package main

import (
	"fmt"
	"github.com/phamhongviet/godsv"
)

func main() {
	row := godsv.Row{
		"this",
		"is",
		"a:simple",
		"test",
		"te\\st",
	}
	fmt.Println(godsv.Marshal(row))
}
