package godsv

import (
	"fmt"
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

func TestUnmarshal(t *testing.T) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	expectedResult := Row{
		"abc",
		"def",
		"ghk",
		"",
		"ab:cd",
		"ef\\gh",
	}

	result := Unmarshal(sample)
	for k, v := range expectedResult {
		if result[k] != v {
			t.Errorf("Unmarshal failed.\nWant: %s\nGot:  %s", expectedResult[k], result[k])
			return
		}
	}
	return
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
