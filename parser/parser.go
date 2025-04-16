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
		return il.Operator + " " + il.Right.String()
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

type RecursiveDescentParser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
}

func New(l *lexer.Lexer) Parser {
	return &RecursiveDescentParser{
		l:            l,
		currentToken: l.NextToken(),
	}
}

func (p *RecursiveDescentParser) ParseProgram() *Program {
	return &Program{
		Expression: p.parseAddtion(),
	}
}

func (p *RecursiveDescentParser) parseAddtion() Expression {
	left := p.parseMultiplication()
	for p.currentToken.Type == lexer.PLUS || p.currentToken.Type == lexer.MINUS {
		operator := p.currentToken.Literal
		p.nextToken()
		right := p.parseMultiplication()
		left = &InfixExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

func (p *RecursiveDescentParser) parseMultiplication() Expression {
	left := p.getIntegerLiteralAsExpression()
	p.nextToken()
	for p.currentToken.Type == lexer.MULTI || p.currentToken.Type == lexer.DIVIDE {
		operator := p.currentToken.Literal
		p.nextToken()
		right := p.getIntegerLiteralAsExpression()
		p.nextToken()
		left = &InfixExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

func (p *RecursiveDescentParser) nextToken() {
	t := p.l.NextToken()
	p.currentToken = t
}

func (p *RecursiveDescentParser) getIntegerLiteralAsExpression() Expression {
	switch p.currentToken.Type {
	case lexer.INT:
		return getIntegerLiteral(p.currentToken.Literal)
	case lexer.PLUS, lexer.MINUS:
		operator := p.currentToken.Literal
		p.nextToken()
		right := p.getIntegerLiteralAsExpression()
		return &InfixExpression{
			Operator: operator,
			Right:    right,
		}
	case lexer.LPAREN:
		p.nextToken()
		expr := p.parseAddtion()
		if p.currentToken.Type != lexer.RPAREN {
			panic("右括弧が見つかりません。")
		}
		return expr
	default:
		panic("paser.getIntegerLiteralAsExpression()が意図しない動作をしています。")
	}
}

func getIntegerLiteral(literal string) *IntegerLiteral {
	value, err := strconv.ParseInt(literal, 10, 64)
	if err != nil {
		panic("strconv.ParseInt()でエラーが発生しました。")
	}
	return &IntegerLiteral{Value: value}
}
