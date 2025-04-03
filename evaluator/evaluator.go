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
	default:
		return nil
	}
}
