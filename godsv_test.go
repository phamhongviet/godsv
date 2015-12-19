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
	dsv := New()
	fmt.Println(dsv.Marshal(sample))
	fmt.Println(strings.Join(sample, " "))
	// Output: abc:def:ghk::ab\:cd:ef\\gh
	// abc def ghk  ab:cd ef\gh
}

func ExampleUnmarshal() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	dsv := New()
	result := dsv.Unmarshal(sample)
	fmt.Println(strings.Join(result, " "))
	fmt.Println(sample)
	// Output: abc def ghk  ab:cd ef\gh
	// abc:def:ghk::ab\:cd:ef\\gh
}

func ExampleCustomMarshal() {
	sample := Row{
		"abc",
		"def",
		"ghk",
		"",
		"ab,cd",
		"ef\\gh",
	}
	dsv := NewCustom(',', defaultEscape)
	fmt.Println(dsv.Marshal(sample))
	fmt.Println(strings.Join(sample, " "))
	// Output: abc,def,ghk,,ab\,cd,ef\\gh
	// abc def ghk  ab,cd ef\gh
}

func ExampleCustomUnmarshal() {
	sample := "abc,def,ghk,,ab\\,cd,ef\\\\gh"
	dsv := NewCustom(',', defaultEscape)
	result := dsv.Unmarshal(sample)
	fmt.Println(strings.Join(result, " "))
	fmt.Println(sample)
	// Output: abc def ghk  ab,cd ef\gh
	// abc,def,ghk,,ab\,cd,ef\\gh
}

func ExampleCut() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	dsv := New()
	result, leftover := dsv.cut(sample)
	fmt.Println(result)
	fmt.Println(leftover)
	// Output: abc
	// def:ghk::ab\:cd:ef\\gh
}

func ExampleCount() {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	dsv := New()
	fmt.Printf("%d\n", dsv.count(sample))
	fmt.Println(sample)
	// Output: 6
	// abc:def:ghk::ab\:cd:ef\\gh
}

func ExampleClean() {
	sample := "ab\\:cd\\\\e"
	dsv := New()
	fmt.Println(dsv.clean(sample))
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
	dsv := New()
	for i := 0; i < b.N; i++ {
		dsv.Marshal(sample)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	dsv := New()
	for i := 0; i < b.N; i++ {
		dsv.Unmarshal(sample)
	}
}

func BenchmarkCount(b *testing.B) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	dsv := New()
	for i := 0; i < b.N; i++ {
		dsv.count(sample)
	}
}

func BenchmarkClean(b *testing.B) {
	sample := "ab\\:cd\\\\e"
	dsv := New()
	for i := 0; i < b.N; i++ {
		dsv.clean(sample)
	}
}
