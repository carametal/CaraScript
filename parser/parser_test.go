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
