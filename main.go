package main

import (
	"fmt"
	"github.com/baxromumarov/calculator-go/parser"
)

func main() {
	text := "if 0==0 then ((((1 + 2) * 3) - 4) / 5) * 6 "

	result := parser.Calculator(text)
	fmt.Println("result: ", result)
}
