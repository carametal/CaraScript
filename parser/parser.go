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

type NumberLiteral struct {
	Value float64
}

func (i *NumberLiteral) expresssionNode() {

}

func (i *NumberLiteral) String() string {
	return strconv.FormatFloat(i.Value, 'f', -1, 64)
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
		Expression: p.parseAddition(),
	}
}

func (p *RecursiveDescentParser) parseAddition() Expression {
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
	left := p.getNumberLiteralAsExpression()
	p.nextToken()
	for p.currentToken.Type == lexer.MULTI || p.currentToken.Type == lexer.DIVIDE {
		operator := p.currentToken.Literal
		p.nextToken()
		right := p.getNumberLiteralAsExpression()
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

func (p *RecursiveDescentParser) getNumberLiteralAsExpression() Expression {
	switch p.currentToken.Type {
	case lexer.INT, lexer.FLOAT:
		return getNumberLiteral(p.currentToken.Literal)
	case lexer.PLUS, lexer.MINUS:
		operator := p.currentToken.Literal
		p.nextToken()
		right := p.getNumberLiteralAsExpression()
		return &InfixExpression{
			Operator: operator,
			Right:    right,
		}
	case lexer.LPAREN:
		p.nextToken()
		expr := p.parseAddition()
		if p.currentToken.Type != lexer.RPAREN {
			panic("右括弧が見つかりません。")
		}
		return expr
	default:
		panic("paser.getNumberLiteralAsExpression()が意図しない動作をしています。")
	}
}

func getNumberLiteral(literal string) *NumberLiteral {
	value, err := strconv.ParseFloat(literal, 64)
	if err != nil {
		panic("strconv.ParseFloat()でエラーが発生しました。")
	}
	return &NumberLiteral{Value: value}
}
