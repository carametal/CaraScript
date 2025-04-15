package parser

import (
	"carametal/CaraScript/lexer"
	"testing"
)

func TestNewSimpleParser(t *testing.T) {
	l := lexer.New("1")
	p := NewSimpleParser(l)
	if p == nil {
		t.Fatal("Parserのインスタンス化に失敗しました。")
	}
}

func TestParseProgram(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "単一の整数",
			input: "123",
			want:  "123",
		},
		{
			name:  "単一のプラストークン",
			input: "+",
			want:  "+",
		},
		{
			name:  "2つの整数の足し算",
			input: "123+456",
			want:  "123 + 456",
		},
		{
			name:  "単一のマイナストークン",
			input: "-",
			want:  "-",
		},
		{
			name:  "2つの整数の引き算",
			input: "123-24",
			want:  "123 - 24",
		},
		{
			name:  "単一のマルチトークン",
			input: "*",
			want:  "*",
		},
		{
			name:  "2つの整数の掛け算",
			input: "123*24",
			want:  "123 * 24",
		},
		{
			name:  "単一のディバイドトークン",
			input: "/",
			want:  "/",
		},
		{
			name:  "2つの正の整数の割り算",
			input: "123/45",
			want:  "123 / 45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := NewSimpleParser(l)
			program := p.ParseProgram()
			if program.Expression.String() != tt.want {
				t.Fatalf("program.Expression want %s, got %s", tt.want, program.Expression)
			}
		})
	}
}

func TestNewRecursiveDescentParser(t *testing.T) {
	l := lexer.New("123")
	p := NewRecursiveDescentParser(l)
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
			name:  "単一の整数",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := NewRecursiveDescentParser(l)
			program := p.ParseProgram()
			if program.Expression.String() != tt.want {
				t.Fatalf("program.Expression want %s, got %s", tt.want, program.Expression)
			}
		})
	}
}
