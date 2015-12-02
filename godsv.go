package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const Delimiter = ":"
const DelimiterRune = ':'

// Escape is the character use for escaping delimiter in values
const Escape = "\\"
const EscapeRune = '\\'

// Marshal encode a Row into a line in DSV file
func Marshal(row Row) string {
	for k, v := range row {
		// escape special characters (Delimiter and Escape itself)
		v = strings.Replace(v, Escape, Escape+Escape, -1)
		row[k] = strings.Replace(v, Delimiter, Escape+Delimiter, -1)
	}
	return strings.Join(row, Delimiter)
}

// Unmarshal decode a line in DSV file into a Row
func Unmarshal(line string) Row {
	return nil
}

// count delimiters in line
func count(line string) int {
	result := 1

	literal := false
	for _, v := range line {
		switch {
		case literal:
			{
				literal = false
				continue
			}
		case v == EscapeRune:
			{
				literal = true
				continue
			}
		case v == DelimiterRune:
			{
				result += 1
			}
		}
	}

	return result
}
