package calc

import (
	"bufio"
	"strings"
)

type Calculator struct {
	operands []string
	result   int
	scn      *bufio.Scanner
}

// New creates a Calculator
func New(scn *bufio.Scanner) Calculator {
	return Calculator{operands: []string{}, result: 0, scn: scn}
}

// rmSpaces remove space separator from calculation lines.
func rmSpaces(strs, fields []string) []string {
	if len(strs) == 0 {
		return fields
	}
	s := strings.Fields(strs[0])
	fields = append(fields, s...)
	return rmSpaces(strs[1:], fields)
}
