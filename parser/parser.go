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

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (il *InfixExpression) expresssionNode() {}

func (il *InfixExpression) String() string {
	if il.Left == nil && il.Right == nil {
		return il.Operator
	}
	if il.Left == nil {
		return il.Right.String()
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

type Parser interface {
	ParseProgram() *Program
}

type SimpleParser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
}

func NewSimpleParser(l *lexer.Lexer) Parser {
	return &SimpleParser{
		l:            l,
		currentToken: l.NextToken(),
	}
}

func (p *SimpleParser) ParseProgram() *Program {
	program := &Program{}
	for p.currentToken.Type != lexer.EOF {
		switch p.currentToken.Type {
		case lexer.INT:
			exp, isInfixLiteral := program.Expression.(*InfixExpression)
			if isInfixLiteral {
				exp.Right = getIntegerLiteral(p.currentToken.Literal)
			} else {
				program.Expression = getIntegerLiteral(p.currentToken.Literal)
			}
		case lexer.PLUS, lexer.MINUS, lexer.MULTI, lexer.DIVIDE:
			// case lexer.PLUS:
			left, isLeftInteger := program.Expression.(*IntegerLiteral)
			operator := p.currentToken.Literal
			if !isLeftInteger {
				program.Expression = &InfixExpression{
					Operator: operator,
				}
			} else {
				program.Expression = &InfixExpression{
					Left:     left,
					Operator: operator,
				}
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

type RecursiveDescentParser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
}

func NewRecursiveDescentParser(l *lexer.Lexer) Parser {
	return &RecursiveDescentParser{
		l:            l,
		currentToken: l.NextToken(),
	}
}

func (p *RecursiveDescentParser) ParseProgram() *Program {
	return nil
}
