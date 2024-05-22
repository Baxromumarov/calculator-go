package parser

import (
	"fmt"
	tk "github.com/baxromumarov/calculator-go/token"
	"strconv"
)




// Parser represents a parser for parsing tokens.
type Parser struct {
	lexer  *tk.Lexer
	tokens []tk.Token
	pos    int
}

// NewParser initializes a new Parser.
func newParser(lexer *tk.Lexer) *Parser {
	return &Parser{lexer: lexer, pos: 0}
}

// eat consumes a token of a given type.
func (p *Parser) eat(tokenType tk.TokenType) {
	if p.currentToken().Type == tokenType {
		p.pos++
	} else {
		panic(fmt.Sprintf("Unexpected token: %v", p.currentToken()))
	}
}

// currentToken returns the current token.
func (p *Parser) currentToken() tk.Token {
	return p.tokens[p.pos]
}

// factor parses and returns a factor.
func (p *Parser) factor() float64 {
	token := p.currentToken()
	if token.Type == tk.NUMBER {
		p.eat(tk.NUMBER)
		value, _ := strconv.ParseFloat(token.Value, 64)
		return value
	} else if token.Type == tk.LPAREN {
		p.eat(tk.LPAREN)
		result := p.expr()
		p.eat(tk.RPAREN)
		return result
	}
	panic(fmt.Sprintf("Unexpected token in factor: %v", token))
}

// term parses and returns a term.
func (p *Parser) term() float64 {
	result := p.factor()
	for p.currentToken().Type == tk.MUL || p.currentToken().Type == tk.DIV {
		token := p.currentToken()
		if token.Type == tk.MUL {
			p.eat(tk.MUL)
			result *= p.factor()
		} else if token.Type == tk.DIV {
			p.eat(tk.DIV)
			result /= p.factor()
		}
	}
	return result
}

// expr parses and returns an expression.
func (p *Parser) expr() float64 {
	result := p.term()
	for p.currentToken().Type == tk.PLUS || p.currentToken().Type == tk.MINUS {
		token := p.currentToken()
		if token.Type == tk.PLUS {
			p.eat(tk.PLUS)
			result += p.term()
		} else if token.Type == tk.MINUS {
			p.eat(tk.MINUS)
			result -= p.term()
		}
	}
	return result
}

// comparison parses and returns a comparison.
func (p *Parser) comparison() bool {
	left := p.expr()
	token := p.currentToken()
	if token.Type == tk.LT {
		p.eat(tk.LT)
		return left < p.expr()
	} else if token.Type == tk.GT {
		p.eat(tk.GT)
		return left > p.expr()
	} else if token.Type == tk.LE {
		p.eat(tk.LE)
		return left <= p.expr()
	} else if token.Type == tk.GE {
		p.eat(tk.GE)
		return left >= p.expr()
	} else if token.Type == tk.EQ {
		p.eat(tk.EQ)
		return left == p.expr()
	} else if token.Type == tk.NE {
		p.eat(tk.NE)
		return left != p.expr()
	}
	panic(fmt.Sprintf("Unexpected token in comparison: %v", token))
}

// statement parses and executes a statement.
func (p *Parser) statement() float64 {
	if p.currentToken().Type == tk.IF {
		p.eat(tk.IF)
		cond := p.comparison()
		p.eat(tk.THEN)
		if cond {
			result := p.expr()
			if p.currentToken().Type == tk.ELSE {
				p.eat(tk.ELSE)
				p.expr() // Consume else part
			}
			return result
		} else {
			if p.currentToken().Type == tk.ELSE {
				p.eat(tk.ELSE)
				return p.expr()
			}
		}
	}
	return p.expr()
}

// parse initializes the token list and parses the input.
func (p *Parser) parse() float64 {
	for {
		token := p.lexer.GetNextToken()
		p.tokens = append(p.tokens, token)
		if token.Type == tk.EOF {
			break
		}
	}
	return p.statement()
}

func Calculator(expression string) float64 {
	lexer := tk.NewLexer(expression)
	parser := newParser(lexer)
	result := parser.parse()
	//fmt.Printf("Result: %f\n", result)
	return result
}
