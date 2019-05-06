// Package number provides functions to obtain formatted numbers as strings.
package number

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// CommaFormat returns the string number `s` formatted with commas.
// Decimal points will be left as it is, if any.
func CommaFormat(s string) (string, error) {
	var buf bytes.Buffer

	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return "", fmt.Errorf("Format: %v", err)
	}

	n := len(s)
	dot := strings.LastIndex(s, ".")
	if dot > 0 {
		n = dot
	}

	if n <= 3 {
		return s, nil
	}

	start := 0
	if s[0] == '-' || s[0] == '+' {
		start = 1
		buf.WriteByte(s[0])
	}

	for i := start; i < n; i++ {
		if i > start && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	if dot > 0 {
		buf.WriteString(s[dot:])
	}

	return buf.String(), nil
}
