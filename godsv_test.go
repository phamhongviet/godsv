package godsv

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	sample := Row{
		"abc",
		"def",
		"ghk",
		"",
		"ab:cd",
		"ef\\gh",
	}
	expectedResult := "abc:def:ghk::ab\\:cd:ef\\\\gh"

	result := Marshal(sample)
	if expectedResult != result {
		t.Errorf("Marshal failed.\nWant: %s\nGot: %s", expectedResult, result)
	}
	return
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

func TestCount(t *testing.T) {
	sample := "abc:def:ghk::ab\\:cd:ef\\\\gh"
	expectedResult := 6

	result := count(sample)
	if result != expectedResult {
		t.Errorf("Count failed.\nWant: %d\nGot:  %d", expectedResult, result)
	}
}

func TestClean(t *testing.T) {
	sample := "ab\\:cd\\\\e"
	expectedResult := "ab:cd\\e"

	result := clean(sample)
	if result != expectedResult {
		t.Errorf("Clean failed.\nWant: %s\nGot:  %s", expectedResult, result)
	}
}
