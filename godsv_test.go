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
	expectedResult := "abc:def:ghk::ab\\:cd:ef\\gh"

	if expectedResult != Marshal(sample) {
		t.Errorf("Marshal failed")
	}
	return
}

func TestUnmarshal(t *testing.T) {
	sample := "abc:def:ghk::ab\\:cd:ef\\gh"
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
			t.Errorf("Unmarshal failed")
		}
	}
	return
}
