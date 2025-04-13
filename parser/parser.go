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

type InfixLiteral struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (il *InfixLiteral) expresssionNode() {}

func (il *InfixLiteral) String() string {
	if il.Left == nil && il.Right == nil {
		return il.Operator
	}
	return il.Left.String() + " " + il.Operator + " " + il.Right.String()
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
	program := &Program{}
	for p.currentToken.Type != lexer.EOF {
		switch p.currentToken.Type {
		case lexer.INT:
			exp, isInfixLiteral := program.Expression.(*InfixLiteral)
			if isInfixLiteral {
				exp.Right = getIntegerLiteral(p.currentToken.Literal)
			} else {
				program.Expression = getIntegerLiteral(p.currentToken.Literal)
			}
		case lexer.PLUS:
			left := program.Expression
			operator := p.currentToken.Literal
			program.Expression = &InfixLiteral{
				Left:     left,
				Operator: operator,
				Right:    nil,
			}
		}
		p.currentToken = p.l.NextToken()
	}
	return program
}

func getIntegerLiteral(literal string) *IntegerLiteral {
	value, err := strconv.ParseInt(literal, 10, 64)
	if err != nil {
		panic("strconv.ParseInt()でエラーが発生しました。")
	}
	return &IntegerLiteral{Value: value}
}
