package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

type DSVParser struct {
	delimiterRune    rune
	escapeRune       rune
	escape           string
	escapedEscape    string
	delimiter        string
	escapedDelimiter string
}

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const defaultDelimiter = ':'

// Escape is the character use for escaping delimiter in values
const defaultEscape = '\\'

func New() DSVParser {
	return NewCustom(defaultDelimiter, defaultEscape)
}

func NewCustom(delimiterRune rune, escapeRune rune) DSVParser {
	return DSVParser{
		delimiterRune:    delimiterRune,
		escapeRune:       escapeRune,
		escape:           string(escapeRune),
		escapedEscape:    string(escapeRune) + string(escapeRune),
		delimiter:        string(delimiterRune),
		escapedDelimiter: string(escapeRune) + string(delimiterRune),
	}
}

// Marshal encode a Row into a line in DSV file
func (dsv DSVParser) Marshal(row Row) string {
	tempRow := make(Row, len(row))

	for k, v := range row {
		tempRow[k] = v
		// escape special characters (Delimiter and Escape itself)
		tempRow[k] = strings.Replace(v, dsv.escape, dsv.escapedEscape, -1)
		tempRow[k] = strings.Replace(tempRow[k], dsv.delimiter, dsv.escapedDelimiter, -1)
	}
	return strings.Join(tempRow, dsv.delimiter)
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
		case v == dsv.escapeRune:
			literal = true
		case v == dsv.delimiterRune:
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
		case v == dsv.escapeRune:
			literal = true
			continue
		case v == dsv.delimiterRune:
			result++
		default:
			continue
		}
	}

	return result
}

// clean escape out of a row's value
func (dsv DSVParser) clean(value string) string {
	result := strings.Replace(value, dsv.escapedDelimiter, dsv.delimiter, -1)
	result = strings.Replace(result, dsv.escapedEscape, dsv.escape, -1)
	return string(result)
}
