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
	fmt.Println(strings.Join(sample, " "))
	// Output: abc:def:ghk::ab\:cd:ef\\gh
	// abc def ghk  ab:cd ef\gh
}

func ExampleUnmarshal() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	result := Unmarshal(sample)
	fmt.Println(strings.Join(result, " "))
	fmt.Println(sample)
	// Output: abc def ghk  ab:cd ef\gh
	// abc:def:ghk::ab\:cd:ef\\gh
}

func ExampleCut() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	result, leftover := cut(sample)
	fmt.Println(result)
	fmt.Println(leftover)
	// Output: abc
	// def:ghk::ab\:cd:ef\\gh
}

func ExampleCount() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	fmt.Printf("%d\n", count(sample))
	fmt.Println(sample)
	// Output: 6
	// abc:def:ghk::ab\:cd:ef\\gh
}

func ExampleClean() {
	sample := "ab\\:cd\\\\e"
	fmt.Println(clean(sample))
	fmt.Println(sample)
	// Output: ab:cd\e
	// ab\:cd\\e
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

func BenchmarkCount(b *testing.B) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	for i := 0; i < b.N; i++ {
		count(sample)
	}
}

func BenchmarkClean(b *testing.B) {
	sample := "ab\\:cd\\\\e"
	for i := 0; i < b.N; i++ {
		clean(sample)
	}
}
