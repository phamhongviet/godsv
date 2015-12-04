package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const Delimiter = ":"

// DelimiterRune is a delimiter as a rune
const DelimiterRune = ':'

// Escape is the character use for escaping delimiter in values
const Escape = "\\"

// EscapeRune is an escape as a rune
const EscapeRune = '\\'

// Marshal encode a Row into a line in DSV file
func Marshal(row Row) string {
	tempRow := make(Row, len(row))
	for k, v := range row {
		tempRow[k] = v
		// escape special characters (Delimiter and Escape itself)
		tempRow[k] = strings.Replace(v, Escape, Escape+Escape, -1)
		tempRow[k] = strings.Replace(tempRow[k], Delimiter, Escape+Delimiter, -1)
	}
	return strings.Join(tempRow, Delimiter)
}

// Unmarshal decode a line in DSV file into a Row
func Unmarshal(line string) Row {
	size := count(line)
	row := make(Row, size)

	for k := range row {
		row[k], line = cut(line)
	}

	return row
}

// cut the first value out of a line
func cut(line string) (value string, leftover string) {
	literal := false
	done := false
	ve := len(line)
	for k, v := range line {
		switch {
		case literal:
			literal = false
		case v == EscapeRune:
			literal = true
		case v == DelimiterRune:
			leftover = line[k+1:]
			ve = k
			done = true
		default:
			continue
		}
		if done {
			break
		}
	}
	value = clean(line[0:ve])
	return value, leftover
}

// count delimiters in line
func count(line string) int {
	result := 1

	literal := false
	for _, v := range line {
		switch {
		case literal:
			literal = false
			continue
		case v == EscapeRune:
			literal = true
			continue
		case v == DelimiterRune:
			result++
		default:
			continue
		}
	}

	return result
}

// clean escape out of a row's value
func clean(value string) string {
	result := strings.Replace(value, Escape+Delimiter, Delimiter, -1)
	result = strings.Replace(result, Escape+Escape, Escape, -1)
	return string(result)
}
