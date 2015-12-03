package godsv

import (
	"fmt"
	"strings"
)

func ExampleMarshal() {
	sample := Row{
		"abc",
		"def",
		"ghk",
		"",
		"ab:cd",
		"ef\\gh",
	}
	fmt.Println(Marshal(sample))
	// Output: abc:def:ghk::ab\:cd:ef\\gh
}

func ExampleUnmarshal() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	result := Unmarshal(sample)
	fmt.Println(strings.Join(result, " "))
	// Output: abc def ghk  ab:cd ef\gh
}

func ExampleCount() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	fmt.Printf("%d\n", count(sample))
	// Output: 6
}

func ExampleClean() {
	sample := "ab\\:cd\\\\e"
	fmt.Println(clean(sample))
	// Output: ab:cd\e
}
