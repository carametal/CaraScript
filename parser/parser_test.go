package parser

import (
	"carametal/CaraScript/lexer"
	"testing"
)

func TestNewRecursiveDescentParser(t *testing.T) {
	l := lexer.New("123")
	p := New(l)
	if p == nil {
		t.Fatal("RecursiveDescentParserのインスタンス化に失敗しました。")
	}
	_, ok := p.(*RecursiveDescentParser)
	if !ok {
		t.Fatal("NewRecursiveDescentParserの結果がRecursiveDescentParserではありません。")
	}
}

func TestParseProgram_RecursiveDescentParser(t *testing.T) {
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
			name:  "複数の正の整数の足し算",
			input: "123+456+789",
			want:  "123 + 456 + 789",
		},
		{
			name:  "複数の正の整数の引き算",
			input: "987-654-321",
			want:  "987 - 654 - 321",
		},
		{
			name:  "複数の正の整数の掛け算",
			input: "123*456*789",
			want:  "123 * 456 * 789",
		},
		{
			name:  "複数の正の整数の割り算",
			input: "123/456/789",
			want:  "123 / 456 / 789",
		},
		{
			name:  "複数の整数による混在した四則演算",
			input: "12+34*56-78/90",
			want:  "12 + 34 * 56 - 78 / 90",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			program := p.ParseProgram()
			if program.Expression.String() != tt.want {
				t.Fatalf("program.Expression want %s, got %s", tt.want, program.Expression)
			}
		})
	}
}
