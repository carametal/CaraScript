package evaluator

import (
	"carametal/CaraScript/parser"
	"strconv"
)

type Object interface {
	String() string
}

type Number struct {
	Value float64
}

func (i *Number) String() string {
	return strconv.FormatFloat(i.Value, 'f', -1, 64)
}

func Eval(node parser.Node) Object {
	switch n := node.(type) {
	case *parser.Program:
		return Eval(n.Expression)
	case *parser.NumberLiteral:
		return &Number{Value: n.Value}
	case *parser.InfixExpression:
		return evalInfixExpression(n)
	default:
		return nil
	}
}

func evalInfixExpression(il *parser.InfixExpression) Object {
	var l, r float64
	if il.Left != nil {
		switch ill := il.Left.(type) {
		case *parser.InfixExpression:
			l = evalInfixExpression(ill).(*Number).Value
		case *parser.NumberLiteral:
			l = ill.Value
		default:
			panic("il.Leftが意図しない値です。il.Left=" + ill.String())
		}
	} else {
		l = 0
	}

	if il.Right != nil {
		switch ilr := il.Right.(type) {
		case *parser.InfixExpression:
			r = evalInfixExpression(ilr).(*Number).Value
		case *parser.NumberLiteral:
			r = ilr.Value
		default:
			panic("il.Rightが意図しない値です。il.Right=" + ilr.String())
		}
	} else {
		r = 0
	}

	switch il.Operator {
	case "+":
		return &Number{
			Value: l + r,
		}
	case "-":
		return &Number{
			Value: l - r,
		}
	case "*":
		return &Number{
			Value: l * r,
		}
	case "/":
		if r == 0 {
			panic("0除算が発生しました")
		}
		return &Number{
			Value: l / r,
		}
	default:
		panic("Eval()で意図しない挙動が発生しています。")
	}
}
