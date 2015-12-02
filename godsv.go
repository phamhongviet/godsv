package godsv

import (
	"strings"
)

// Row represent a row in DSV file
type Row []string

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const Delimiter = ":"

// Escape is the character use for escaping delimiter in values
const Escape = "\\"

// Marshal encode a Row into a line in DSV file
func Marshal(row Row) (line string) {
	for _, v := range row {
		// escape special characters (Delimiter and Escape itself)
		v = strings.Replace(v, Escape, Escape+Escape, -1)
		v = strings.Replace(v, Delimiter, Escape+Delimiter, -1)
		line += v + Delimiter
	}
	return strings.TrimSuffix(line, Delimiter)
}

// Unmarshal decode a line in DSV file into a Row
func Unmarshal(line string) Row {
	return nil
}
