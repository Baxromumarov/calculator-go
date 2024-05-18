package token

import (
	"fmt"
	"unicode"
)

// TokenType represents the type of token.
type TokenType string

const (
	NUMBER TokenType = "NUMBER"
	PLUS   TokenType = "PLUS"
	MINUS  TokenType = "MINUS"
	MUL    TokenType = "MUL"
	DIV    TokenType = "DIV"
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	IF     TokenType = "IF"
	THEN   TokenType = "THEN"
	ELSE   TokenType = "ELSE"
	LT     TokenType = "LT"
	GT     TokenType = "GT"
	LE     TokenType = "LE"
	GE     TokenType = "GE"
	EQ     TokenType = "EQ"
	NE     TokenType = "NE"
	EOF    TokenType = "EOF"
	IDENT  TokenType = "IDENT"
)

// Token represents a token with type and value.
type Token struct {
	Type  TokenType
	Value string
}

// Lexer represents a lexer for tokenizing input strings.
type Lexer struct {
	text        string
	pos         int
	currentChar rune
}

// NewLexer initializes a new Lexer.
func NewLexer(text string) *Lexer {
	l := &Lexer{text: text, pos: 0}
	l.currentChar = rune(text[l.pos])
	return l
}

// advance moves the lexer position one step forward.
func (l *Lexer) advance() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0 // EOF
	} else {
		l.currentChar = rune(l.text[l.pos])
	}
}

// skipWhitespace skips any whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) {
		l.advance()
	}
}

// identifier handles identifiers and keywords.
func (l *Lexer) identifier() string {
	result := ""
	for l.currentChar != 0 && (unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar)) {
		result += string(l.currentChar)
		l.advance()
	}
	return result
}

// number handles multi-digit numbers.
func (l *Lexer) number() string {
	result := ""
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) {
		result += string(l.currentChar)
		l.advance()
	}
	return result
}

// GetNextToken generates the next token from the input.
func (l *Lexer) GetNextToken() Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.skipWhitespace()
			continue
		}
		if unicode.IsLetter(l.currentChar) {
			value := l.identifier()
			switch value {
			case "if":
				return Token{Type: IF, Value: value}
			case "then":
				return Token{Type: THEN, Value: value}
			case "else":
				return Token{Type: ELSE, Value: value}
			default:
				return Token{Type: IDENT, Value: value}
			}
		}
		if unicode.IsDigit(l.currentChar) {
			return Token{Type: NUMBER, Value: l.number()}
		}
		if l.currentChar == '+' {
			l.advance()
			return Token{Type: PLUS, Value: "+"}
		}
		if l.currentChar == '-' {
			l.advance()
			return Token{Type: MINUS, Value: "-"}
		}
		if l.currentChar == '*' {
			l.advance()
			return Token{Type: MUL, Value: "*"}
		}
		if l.currentChar == '/' {
			l.advance()
			return Token{Type: DIV, Value: "/"}
		}
		if l.currentChar == '(' {
			l.advance()
			return Token{Type: LPAREN, Value: "("}
		}
		if l.currentChar == ')' {
			l.advance()
			return Token{Type: RPAREN, Value: ")"}
		}
		if l.currentChar == '<' {
			l.advance()
			if l.currentChar == '=' {
				l.advance()
				return Token{Type: LE, Value: "<="}
			}
			return Token{Type: LT, Value: "<"}
		}
		if l.currentChar == '>' {
			l.advance()
			if l.currentChar == '=' {
				l.advance()
				return Token{Type: GE, Value: ">="}
			}
			return Token{Type: GT, Value: ">"}
		}
		if l.currentChar == '=' {
			l.advance()
			if l.currentChar == '=' {
				l.advance()
				return Token{Type: EQ, Value: "=="}
			}
		}
		if l.currentChar == '!' {
			l.advance()
			if l.currentChar == '=' {
				l.advance()
				return Token{Type: NE, Value: "!="}
			}
		}
		panic(fmt.Sprintf("Unknown character: %v", l.currentChar))
	}
	return Token{Type: EOF, Value: ""}
}
