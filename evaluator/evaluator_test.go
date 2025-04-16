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
			name:  "3つの正の整数の足し算",
			input: "123+456+789",
			want:  "1368",
		},
		{
			name:  "3つの正の整数の引き算",
			input: "789-456-123",
			want:  "210",
		},
		{
			name:  "差が負の整数になる2つの正の整数の引き算",
			input: "123-456",
			want:  "-333",
		},
		{
			name:  "3つの正の整数の掛け算",
			input: "123*456*789",
			want:  "44253432",
		},
		{
			name:  "割り切れない2つの正の整数の割り算",
			input: "123/45",
			want:  "2",
		},
		{
			name:  "3つの正の整数の割り算",
			input: "987/6/5",
			want:  "32",
		},
		{
			name:  "計算すべき順序が決まっている複数の整数の四則演算",
			input: "1+2*3",
			want:  "7",
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
