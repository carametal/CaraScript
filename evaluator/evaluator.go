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
	switch n := node.(type) {
	case *parser.Program:
		return Eval(n.Expression)
	case *parser.IntegerLiteral:
		return &Integer{Value: n.Value}
	case *parser.InfixExpression:
		return evalInfixExpression(n)
	default:
		return nil
	}
}

func evalInfixExpression(il *parser.InfixExpression) Object {
	var l, r int64
	if il.Left != nil {
		switch ill := il.Left.(type) {
		case *parser.InfixExpression:
			l, _ = strconv.ParseInt(evalInfixExpression(ill).String(), 10, 64)
		case *parser.IntegerLiteral:
			l, _ = strconv.ParseInt(ill.String(), 10, 64)
		default:
			panic("il.Leftが意図しない値です。il.Left=" + ill.String())
		}
	} else {
		l = 0
	}

	if il.Right != nil {
		switch ilr := il.Right.(type) {
		case *parser.InfixExpression:
			r, _ = strconv.ParseInt(evalInfixExpression(ilr).String(), 10, 64)
		case *parser.IntegerLiteral:
			r, _ = strconv.ParseInt(ilr.String(), 10, 64)
		default:
			panic("il.Rightが意図しない値です。il.Right=" + ilr.String())
		}
	} else {
		r = 0
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
		panic("Eval()で意図しない挙動が発生しています。")
	}
}
