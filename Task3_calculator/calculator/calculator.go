package calc

import (
	"bufio"
	"log"
	"strconv"
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

// sum iterates over Calculator.operands and convert each operand to integer
// and return total sum. it will have fatal when faces with first non-integer operand.
func (c *Calculator) sum() int {
	if len(c.operands) == 0 {
		return c.result
	}
	n, err := strconv.Atoi(c.operands[0])
	if err != nil {
		log.Fatalf("syntax error - '%s' not an integer\n", c.operands[0])
	}
	c.result += n
	c.operands = c.operands[1:]
	return c.sum()
}

// readAllCalcLines reads all calculation lines from console.
func (c *Calculator) readAllCalcLines(n int, s []string) []string {
	if n == 0 {
		return s
	}
	if c.scn.Scan() {
		l := c.scn.Text()
		s = append(s, l)
	}
	return c.readAllCalcLines(n-1, s)
}

// rmAllDelimiterSigns removes all delimiter signs from calculation lines.
func rmAllDelimiterSigns(strs []string) []string {
	return rmEmpty(
		rmPlus(
			rmComma(
				rmSpaces(strs, []string{}),
				[]string{}),
			[]string{}),
		[]string{},
	)
}

// rmEmpty remove empty strings from input fields slice.
func rmEmpty(strs, fields []string) []string {
	if len(strs) == 0 {
		return fields
	}
	if strs[0] != "" {
		fields = append(fields, strs[0])
	}
	return rmEmpty(strs[1:], fields)
}

// rmPlus remove plus separator from calculation lines.
func rmPlus(strs, fields []string) []string {
	if len(strs) == 0 {
		return fields
	}
	if strs[0] != "" {
		s := strings.Split(strs[0], "+")
		fields = append(fields, s...)
	}
	return rmPlus(strs[1:], fields)
}

// rmComma remove comma separator from calculation lines.
func rmComma(strs, fields []string) []string {
	if len(strs) == 0 {
		return fields
	}
	if strs[0] != "" {
		s := strings.Split(strs[0], ",")
		fields = append(fields, s...)
	}
	return rmComma(strs[1:], fields)
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
