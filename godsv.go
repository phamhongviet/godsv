package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

type DSVParser struct {
	Delimiter     string
	DelimiterRune rune
	Escape        string
	EscapeRune    rune
}

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const defaultDelimiter = ":"

// DelimiterRune is a delimiter as a rune
const defaultDelimiterRune = ':'

// Escape is the character use for escaping delimiter in values
const defaultEscape = "\\"

// EscapeRune is an escape as a rune
const defaultEscapeRune = '\\'

func New() DSVParser {
	return DSVParser{defaultDelimiter, defaultDelimiterRune, defaultEscape, defaultEscapeRune}
}

// Marshal encode a Row into a line in DSV file
func (dsv DSVParser) Marshal(row Row) string {
	tempRow := make(Row, len(row))
	for k, v := range row {
		tempRow[k] = v
		// escape special characters (Delimiter and Escape itself)
		tempRow[k] = strings.Replace(v, dsv.Escape, dsv.Escape+dsv.Escape, -1)
		tempRow[k] = strings.Replace(tempRow[k], dsv.Delimiter, dsv.Escape+dsv.Delimiter, -1)
	}
	return strings.Join(tempRow, dsv.Delimiter)
}

// Unmarshal decode a line in DSV file into a Row
func (dsv DSVParser) Unmarshal(line string) Row {
	size := dsv.count(line)
	row := make(Row, size)

	for k := range row {
		row[k], line = dsv.cut(line)
	}

	return row
}

// cut the first value out of a line
func (dsv DSVParser) cut(line string) (value string, leftover string) {
	literal := false
	done := false
	ve := len(line)
	for k, v := range line {
		switch {
		case literal:
			literal = false
		case v == dsv.EscapeRune:
			literal = true
		case v == dsv.DelimiterRune:
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
	value = dsv.clean(line[0:ve])
	return value, leftover
}

// count delimiters in line
func (dsv DSVParser) count(line string) int {
	result := 1

	literal := false
	for _, v := range line {
		switch {
		case literal:
			literal = false
			continue
		case v == dsv.EscapeRune:
			literal = true
			continue
		case v == dsv.DelimiterRune:
			result++
		default:
			continue
		}
	}

	return result
}

// clean escape out of a row's value
func (dsv DSVParser) clean(value string) string {
	result := strings.Replace(value, dsv.Escape+dsv.Delimiter, dsv.Delimiter, -1)
	result = strings.Replace(result, dsv.Escape+dsv.Escape, dsv.Escape, -1)
	return string(result)
}
