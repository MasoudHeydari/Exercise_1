package main

import (
	"bufio"
	"os"

	"github.com/MasoudHeydari/Exercise_1/Task2_calculator/calculator"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	c := calc.New(scn)
	c.StarCalculation()
}
