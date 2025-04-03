package lexer

import "strconv"

type TokenType int

const (
	INT = iota
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input           string
	currentPosition int
	peekPosition    int
}

func New(input string) *Lexer {
	return &Lexer{
		input:           input,
		currentPosition: 0,
		peekPosition:    1,
	}
}

func (l *Lexer) NextToken() Token {
	if isDigit(l.input[l.currentPosition]) {
		return Token{
			Literal: l.getDigits(),
			Type:    INT,
		}
	}
	return Token{
		Literal: "",
		Type:    EOF,
	}
}

func (l *Lexer) getDigits() string {
	for len(l.input) > l.peekPosition && isDigit(l.input[l.peekPosition]) {
		l.peekPosition++
	}
	return string(l.input[l.currentPosition:l.peekPosition])
}

func isDigit(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}
