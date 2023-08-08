package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	ADD = iota
	SUB
	MUL
)

const (
	Add = "abcd"
	Sub = "bcde"
	Mul = "dede"
)

var (
	alphabet = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	verbs = map[string]int{
		Add: ADD,
		Sub: SUB,
		Mul: MUL,
	}
)

func main() {
	fmt.Println("starting Xulu lang")
	input := "bcde ab ac" //"abcd abcd aabbc ab a c ccd dede cccd cd" //"bcde ab ac"
	tokens := strings.Fields(input)
	parseTokens(tokens)
}

func parseTokens(tokens []string) {
	s := New()
	g := NewGraph(len(tokens))
	var startNode *Node
	for _, token := range tokens {
		t := Token(token)
		if err := t.IsValid(); err != nil {
			log.Fatal(err)
		}
		var n *Node
		nodeResult := 0
		nodeType := t.GetType()
		if nodeType == TypeVerb {
			if t.IsMUL() {
				nodeResult = 1
			}
		} else {
			nodeResult = t.CalcName()
		}
		n = &Node{Op: t, Type: nodeType, Result: nodeResult}

		if s.Peek() != nil {
			g.AddEdge(s.Peek().(*Node), n)
		} else {
			startNode = n
		}
		s.Push(n)
		s.Peek().(*Node).AddToResult(nodeResult)
		if nodeType == TypeName {
			s.Pop()
		}
	}

	g.DFS(startNode)
}
