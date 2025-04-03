package parser

import (
	"carametal/CaraScript/lexer"
	"strconv"
)

type Node interface {
	String() string
}

type Expression interface {
	Node
	expresssionNode()
}

type IntegerLiteral struct {
	Value int64
}

func (i *IntegerLiteral) expresssionNode() {

}

func (i *IntegerLiteral) String() string {
	return strconv.FormatInt(i.Value, 10)
}

type Program struct {
	Expression Expression
}

func (p *Program) String() string {
	if p.Expression != nil {
		return p.Expression.String()
	}
	return ""
}

type Parser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		l:            l,
		currentToken: l.NextToken(),
	}
}

func (p *Parser) ParseProgram() *Program {
	if p.currentToken.Type == lexer.INT {
		value, err := strconv.ParseInt(p.currentToken.Literal, 10, 64)
		if err != nil {
			return nil
		}
		return &Program{
			Expression: &IntegerLiteral{Value: value},
		}
	}
	return &Program{}
}
