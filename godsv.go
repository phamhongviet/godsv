package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

type DSVParser struct {
	Delimiter    rune
	Escape       rune
}

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const defaultDelimiter = ':'

// Escape is the character use for escaping delimiter in values
const defaultEscape = '\\'

func New() DSVParser {
	return DSVParser{defaultDelimiter, defaultEscape}
}

// Marshal encode a Row into a line in DSV file
func (dsv DSVParser) Marshal(row Row) string {
	tempRow := make(Row, len(row))
	escape := string(dsv.Escape)
	escapedEscape := escape + escape
	delimiter := string(dsv.Delimiter)
	escapedDelimiter := escape + delimiter

	for k, v := range row {
		tempRow[k] = v
		// escape special characters (Delimiter and Escape itself)
		tempRow[k] = strings.Replace(v, escape, escapedEscape, -1)
		tempRow[k] = strings.Replace(tempRow[k], delimiter, escapedDelimiter, -1)
	}
	return strings.Join(tempRow, delimiter)
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
		case v == dsv.Escape:
			literal = true
		case v == dsv.Delimiter:
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
		case v == dsv.Escape:
			literal = true
			continue
		case v == dsv.Delimiter:
			result++
		default:
			continue
		}
	}

	return result
}

// clean escape out of a row's value
func (dsv DSVParser) clean(value string) string {
	escape := string(dsv.Escape)
	escapedEscape := escape + escape
	delimiter := string(dsv.Delimiter)
	escapedDelimiter := escape + delimiter

	result := strings.Replace(value, escapedDelimiter, delimiter, -1)
	result = strings.Replace(result, escapedEscape, escape, -1)
	return string(result)
}
