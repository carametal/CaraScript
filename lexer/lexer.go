package lexer

import (
	"strconv"
)

type TokenType int

const (
	INT = iota
	PLUS
	MINUS
	MULTI
	DIVIDE
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
	l.skipWhitespaces()
	if l.currentPosition >= len(l.input) {
		return Token{
			Literal: "",
			Type:    EOF,
		}
	}
	if isDigit(l.input[l.currentPosition]) {
		return Token{
			Literal: l.getDigits(),
			Type:    INT,
		}
	}
	switch l.input[l.currentPosition] {
	case byte('+'):
		l.currentPosition++
		l.peekPosition++
		return Token{
			Literal: "+",
			Type:    PLUS,
		}
	case byte('-'):
		l.currentPosition++
		l.peekPosition++
		return Token{
			Literal: "-",
			Type:    MINUS,
		}
	case byte('*'):
		l.currentPosition++
		l.peekPosition++
		return Token{
			Literal: "*",
			Type:    MULTI,
		}
	case byte('/'):
		l.currentPosition++
		l.peekPosition++
		return Token{
			Literal: "/",
			Type:    DIVIDE,
		}
	}

	panic("Lexer.NextToken()で予想外の挙動をしています。")
}

func (l *Lexer) getDigits() string {
	for len(l.input) > l.peekPosition && isDigit(l.input[l.peekPosition]) {
		l.peekPosition++
	}
	ret := string(l.input[l.currentPosition:l.peekPosition])
	l.currentPosition = l.peekPosition
	return ret
}

func (l *Lexer) skipWhitespaces() {
	for len(l.input) > l.currentPosition && l.input[l.currentPosition] == byte(' ') {
		l.currentPosition++
		l.peekPosition++
	}
}

func isDigit(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}
