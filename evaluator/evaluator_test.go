package evaluator

import (
	"carametal/CaraScript/lexer"
	"carametal/CaraScript/parser"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "正の整数",
			input: "123",
			want:  "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			got := Eval(p.ParseProgram())
			if got.String() != tt.want {
				t.Fatalf("Eval() want %s, got %s", "1", got.String())
			}
		})
	}
}
