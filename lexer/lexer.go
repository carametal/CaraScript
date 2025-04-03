package lexer

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
	input    string
	position int
}

func New(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
	}
}

func (l *Lexer) NextToken() Token {
	return Token{
		Literal: string(l.input[l.position]),
		Type:    INT,
	}
}
