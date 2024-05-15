package token

import (
	"strings"
)

// Tokenize splits the expression into tokens.
func Tokenize(expression string) []string {
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")
	return strings.Fields(expression)
}
