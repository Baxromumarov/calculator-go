package main

import (
	"fmt"
	"github.com/baxromumarov/calculator-go/parser"
	"github.com/baxromumarov/calculator-go/token"
)

func Calculate(expression string) (float64, error) {
	tokens := token.Tokenize(expression)
	rpn, err := parser.ShuntingYard(tokens)
	if err != nil {
		return 0, err
	}
	return parser.EvaluateRPN(rpn)
}

func main() {
	expression := "2 + 3 * (4 - 1) / 12"
	result, err := Calculate(expression)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of '%s' is: %f\n", expression, result)
	}
}
