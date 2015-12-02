package godsv

// Row represent a row in DSV file
type Row []string

// Delimiter separate values from each other. Delimiter is traditionally a colon, escape with a backslash
const Delimiter = ":"

// Escape is the character use for escaping delimiter in values
const Escape = "\\"

// Marshal encode a Row into a line in DSV file
func Marshal(row Row) string {
	return ""
}

// Unmarshal decode a line in DSV file into a Row
func Unmarshal(line string) Row {
	return nil
}
