package evaluator

import (
	"carametal/CaraScript/parser"
	"math"
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

type Float struct {
	Value float64
}

func (i *Float) String() string {
	return strconv.FormatFloat(i.Value, 'g', 64, 64)
}

func Eval(node parser.Node) Object {
	switch n := node.(type) {
	case *parser.Program:
		return Eval(n.Expression)
	case *parser.IntegerLiteral:
		return &Integer{Value: n.Value}
	case *parser.FloatLiteral:
		return &Float{Value: n.Value}
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
			l, _ = strconv.ParseFloat(evalInfixExpression(ill).String(), 64)
		case *parser.IntegerLiteral:
			l, _ = strconv.ParseFloat(ill.String(), 64)
		default:
			panic("il.Leftが意図しない値です。il.Left=" + ill.String())
		}
	} else {
		l = 0
	}

	if il.Right != nil {
		switch ilr := il.Right.(type) {
		case *parser.InfixExpression:
			r, _ = strconv.ParseFloat(evalInfixExpression(ilr).String(), 64)
		case *parser.IntegerLiteral:
			r, _ = strconv.ParseFloat(ilr.String(), 64)
		default:
			panic("il.Rightが意図しない値です。il.Right=" + ilr.String())
		}
	} else {
		r = 0
	}

	switch il.Operator {
	case "+":
		return makeObject(l + r)
	case "-":
		return makeObject(l - r)
	case "*":
		return makeObject(l * r)
	case "/":
		return makeObject(l / r)
	default:
		panic("Eval()で意図しない挙動が発生しています。")
	}
}

func makeObject(num float64) Object {
	if num == math.Trunc(num) {
		return &Integer{
			Value: int64(num),
		}
	}
	return &Float{
		Value: num,
	}
}
