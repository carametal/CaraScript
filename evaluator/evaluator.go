package evaluator

import (
	"carametal/CaraScript/parser"
	"strconv"
)

type Object interface {
	String() string
}

type Integer struct {
	Value int64
}

func (i *Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
}

func Eval(node parser.Node) Object {
	switch node := node.(type) {
	case *parser.Program:
		return Eval(node.Expression)
	case *parser.IntegerLiteral:
		return &Integer{Value: node.Value}
	case *parser.InfixLiteral:
		return evalInfixLiretal(node)
	default:
		return nil
	}
}

func evalInfixLiretal(il *parser.InfixLiteral) Object {
	switch il.Operator {
	case "+":
		l, _ := strconv.ParseInt(il.Left.String(), 10, 64)
		r, _ := strconv.ParseInt(il.Right.String(), 10, 64)
		return &Integer{
			Value: l + r,
		}
	default:
		panic("Evalで意図しない挙動が発生しています。")
	}
}
