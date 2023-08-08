package main

import (
	"fmt"
	"math"
)

const (
	TypeName = "name"
	TypeVerb = "verb"
)

type Token string

func (t Token) GetType() string {
	if t.isVerb() {
		return TypeVerb
	}
	return TypeName
}

func (t Token) isVerb() bool {
	if _, ok := verbs[t.String()]; ok {
		return true
	}
	return false
}

func (t Token) IsValid() error {
	for _, c := range t {
		currentChar := fmt.Sprintf("%c", c)
		_, ok := alphabet[currentChar]
		if !ok {
			return fmt.Errorf("failed to parse, unknown alphbet('%s') detected", currentChar)
		}
	}
	return nil
}

func (t Token) String() string {
	return string(t)
}

func (t Token) IsMUL() bool {
	return t.String() == "dede"
}

func (t Token) CalcName() int {
	// Split the name into its repeated alphabets
	tokenStr := t.String()
	var repeatedChars []string
	currentChar := string(tokenStr[0])
	for i := 1; i < len(tokenStr); i++ {
		if string(tokenStr[i]) == currentChar[len(currentChar)-1:] {
			currentChar += string(tokenStr[i])
		} else {
			repeatedChars = append(repeatedChars, currentChar)
			currentChar = string(tokenStr[i])
		}
	}
	repeatedChars = append(repeatedChars, currentChar)

	// Compute the mod 5 of each set of repeated alphabets and square them
	var sum int
	for _, al := range repeatedChars {
		i := len(al) * alphabet[string(al[0])] % 5
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

type Node struct {
	Op     Token
	Type   string
	Result int
}

var isFirst = true

func (n *Node) AddToResult(i int) {
	if n.Type == TypeVerb {
		switch n.Op.String() {
		case Mul:
			n.Result *= i
		case Add:
			n.Result += i
		case Sub:
			if isFirst {
				n.Result += i
				isFirst = false
			} else {
				n.Result -= i
			}

		}
	}
}
