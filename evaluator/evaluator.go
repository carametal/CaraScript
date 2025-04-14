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
	case *parser.InfixExpression:
		return evalInfixExpression(node)
	default:
		return nil
	}
}

func evalInfixExpression(il *parser.InfixExpression) Object {
	var l, r int64
	var err error
	if il.Left != nil {
		l, _ = strconv.ParseInt(il.Left.String(), 10, 64)
	}
	r, err = strconv.ParseInt(il.Right.String(), 10, 64)
	if err != nil {
		panic("InfixLiteral.Rightは必須です。")
	}
	switch il.Operator {
	case "+":
		return &Integer{
			Value: l + r,
		}
	case "-":
		return &Integer{
			Value: l - r,
		}
	case "*":
		return &Integer{
			Value: l * r,
		}
	case "/":
		return &Integer{
			Value: l / r,
		}
	default:
		panic("Evalで意図しない挙動が発生しています。")
	}
}
