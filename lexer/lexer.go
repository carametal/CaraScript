package lexer

import (
	"strconv"
	"strings"
)

type TokenType int

const (
	INT TokenType = iota
	FLOAT
	PLUS
	MINUS
	MULTI
	DIVIDE
	LPAREN
	RPAREN
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
	if isDigit(l.input[l.currentPosition]) || (l.input[l.currentPosition] == byte('.') && l.peekPosition < len(l.input) && isDigit(l.input[l.peekPosition])) {
		literal := l.getNumber()
		tType := INT
		if strings.Contains(literal, ".") {
			tType = FLOAT
		}
		return Token{
			Literal: literal,
			Type:    tType,
		}
	}
	switch l.input[l.currentPosition] {
	case byte('+'):
		l.moveNext()
		return Token{
			Literal: "+",
			Type:    PLUS,
		}
	case byte('-'):
		l.moveNext()
		return Token{
			Literal: "-",
			Type:    MINUS,
		}
	case byte('*'):
		l.moveNext()
		return Token{
			Literal: "*",
			Type:    MULTI,
		}
	case byte('/'):
		l.moveNext()
		return Token{
			Literal: "/",
			Type:    DIVIDE,
		}
	case byte('('):
		l.moveNext()
		return Token{
			Literal: "(",
			Type:    LPAREN,
		}
	case byte(')'):
		l.moveNext()
		return Token{
			Literal: ")",
			Type:    RPAREN,
		}
	}

	panic("Lexer.NextToken()で予想外の挙動をしています。")
}

func (l *Lexer) getNumber() string {
	dotFound := false
	for len(l.input) > l.peekPosition && (isDigit(l.input[l.peekPosition]) || (!dotFound && l.input[l.peekPosition] == byte('.'))) {
		if l.input[l.peekPosition] == byte('.') {
			dotFound = true
		}
		l.peekPosition++
	}
	ret := string(l.input[l.currentPosition:l.peekPosition])
	l.currentPosition = l.peekPosition
	return ret
}

func (l *Lexer) skipWhitespaces() {
	for len(l.input) > l.currentPosition {
		switch l.input[l.currentPosition] {
		case byte(' '), '\n', '\t', '\r':
			l.moveNext()
		default:
			return
		}
	}
}

func (l *Lexer) moveNext() {
	l.currentPosition++
	l.peekPosition++
}

func isDigit(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}
