package main

import "testing"

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
	}

	for _, test := range tests {

		result, err := Calculate(test.Input)
		if err != nil {
			t.Error(err)
		}
		if result != test.Expected {
			t.Errorf("Expected %f, got %f", test.Expected, result)
		}
	}

}
