package lexer

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	INT = iota
	PLUS
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
	if l.input[l.currentPosition] == byte('+') {
		l.currentPosition++
		l.peekPosition++
		return Token{
			Literal: "+",
			Type:    PLUS,
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
		fmt.Println("skip", l.input[l.currentPosition])
		l.currentPosition++
		l.peekPosition++
	}
}

func isDigit(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err == nil
}
