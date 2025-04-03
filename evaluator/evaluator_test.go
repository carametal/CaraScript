package evaluator

import (
	"carametal/CaraScript/parser"
	"strconv"
	"testing"
)

func TestEval(t *testing.T) {
	v, _ := strconv.ParseInt("1", 10, 64)
	p := parser.IntegerLiteral{
		Value: v,
	}
	got := Eval(&p)
	if got.String() != "1" {
		t.Fatalf("Eval() want %s, got %s", "1", got.String())
	}
}
