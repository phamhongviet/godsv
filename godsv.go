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
	for k, v := range row {
		// escape special characters (Delimiter and Escape itself)
		v = strings.Replace(v, Escape, Escape+Escape, -1)
		row[k] = strings.Replace(v, Delimiter, Escape+Delimiter, -1)
	}
	return strings.Join(row, Delimiter)
}

// Unmarshal decode a line in DSV file into a Row
func Unmarshal(line string) Row {
	size := count(line)
	row := make(Row, size)

	literal := false
	vi := 0 // value index
	vs := 0 // value start
	ve := 0 // value end
	for k, v := range line {
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
				if ve != 0 {
					vs = ve + 1
				}
				ve = k
				row[vi] = clean(line[vs:ve])
				vi++
			}
		default:
			{
				continue
			}
		}
	}

	if ve != 0 {
		vs = ve + 1
	}
	ve = len(line)
	row[vi] = clean(line[vs:ve])

	return row
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
				result++
			}
		default:
			continue
		}
	}

	return result
}

// clean escape out of a row's value
func clean(value string) string {
	result := ""
	literal := false
	for _, v := range value {
		switch {
		case literal:
			literal = false
			result += string(v)
			continue
		case v == EscapeRune:
			literal = true
			continue
		default:
			result += string(v)
		}
	}
	return result
}
