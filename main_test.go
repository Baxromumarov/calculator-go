package main

import (
	"github.com/baxromumarov/calculator-go/parser"
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		Input    string
		Expected float64
	}{
		{
			"123 + 456 - 789 * 2 / 3",
			53,
		},
		{
			"(10 + 5) * (20 / 4)",
			75,
		},
		{
			"((((1 + 2) * 3) - 4) / 5) * 6",
			6,
		},
		{
			"if 0==0 then ((((1 + 2) * 3) - 4) / 5) * 6 else (10 + 5) * (20 / 4)",
			6,
		},
	}

	for _, test := range tests {

		result := parser.Calculator(test.Input)

		if result != test.Expected {
			t.Errorf("Expected %f, got %f", test.Expected, result)
		}
	}

}
