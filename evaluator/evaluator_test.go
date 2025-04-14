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
			name:  "単一の正の整数",
			input: "123",
			want:  "123",
		},
		{
			name:  "+と単一の整数",
			input: "+123",
			want:  "123",
		},
		{
			name:  "2つの正の整数の足し算",
			input: "123+456",
			want:  "579",
		},
		{
			name:  "スペースを含む2つの正の整数の足し算",
			input: "123 + 456",
			want:  "579",
		},
		{
			name:  "2つの正の整数の引き算",
			input: "123-24",
			want:  "99",
		},
		{
			name:  "スペースを含む2つの正の整数の引き算",
			input: "123 - 24",
			want:  "99",
		},
		{
			name:  "差が負の整数になる2つの正の整数の引き算",
			input: "123-456",
			want:  "-333",
		},
		{
			name:  "2つの正の整数の掛け算",
			input: "123*456",
			want:  "56088",
		},
		{
			name:  "2つの正の整数の割り算",
			input: "124/4",
			want:  "31",
		},
		{
			name:  "割り切れない2つの正の整数の割り算",
			input: "123/45",
			want:  "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)
			got := Eval(p.ParseProgram())
			if got.String() != tt.want {
				t.Fatalf("Eval() want %s, got %s", tt.want, got.String())
			}
		})
	}
}
