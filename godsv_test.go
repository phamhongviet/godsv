package godsv

import (
	"fmt"
	"strings"
	"testing"
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

func BenchmarkMarshal(b *testing.B) {
	sample := Row{
		"abc",
		"def",
		"ghk",
		"",
		"ab:cd",
		"ef\\gh",
	}
	for i := 0; i < b.N; i++ {
		Marshal(sample)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	for i := 0; i < b.N; i++ {
		Unmarshal(sample)
	}
}
