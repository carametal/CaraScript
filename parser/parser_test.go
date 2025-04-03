package parser

import (
	"carametal/CaraScript/lexer"
	"testing"
)

func TestNew(t *testing.T) {
	l := lexer.New("1")
	p := New(l)
	if p == nil {
		t.Fatal("Parserのインスタンス化に失敗しました。")
	}
}

func TestParseProgram(t *testing.T) {
	l := lexer.New("1")
	p := New(l)
	program := p.ParseProgram()
	if program.Expression.String() != "1" {
		t.Fatalf("program.Expression want %s, got %s", "1", program.Expression)
	}
}
