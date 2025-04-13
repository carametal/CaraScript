package lexer

import "testing"

func TestNew(t *testing.T) {
	l := New("1")
	if l == nil {
		t.Fatal("Lexerのインスタンス化に失敗しました。")
	}
	if l.input != "1" {
		t.Fatalf("l.input want %s, got %s", "1", l.input)
	}
	if l.currentPosition != 0 {
		t.Fatalf("l.input want 0, got %d", l.currentPosition)
	}
	if l.peekPosition != 1 {
		t.Fatalf("l.input want 1, got %d", l.peekPosition)
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantTokenType TokenType
		wantLiteral   string
	}{
		{
			name:          "一桁の整数",
			input:         "1",
			wantTokenType: INT,
			wantLiteral:   "1",
		},
		{
			name:          "二桁の整数",
			input:         "12",
			wantTokenType: INT,
			wantLiteral:   "12",
		},
		{
			name:          "単一のプラストークン",
			input:         "+",
			wantTokenType: PLUS,
			wantLiteral:   "+",
		},
		{
			name:          "単一のマイナストークン",
			input:         "-",
			wantTokenType: MINUS,
			wantLiteral:   "-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)
			got := l.NextToken()
			if got.Type != tt.wantTokenType {
				t.Fatalf("l.NextToken().Literal want %d, got %d", tt.wantTokenType, got.Type)
			}
			if got.Literal != tt.wantLiteral {
				t.Fatalf("l.NextToken().Literal want %s, got %s", tt.wantLiteral, got.Literal)
			}
		})
	}
}

func TestGetDisits(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "一桁の整数",
			input: "1",
			want:  "1",
		},
		{
			name:  "二桁の整数",
			input: "12",
			want:  "12",
		},
		{
			name:  "三桁の整数",
			input: "123",
			want:  "123",
		},
	}

	for _, tt := range tests {
		l := New(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			got := l.getDigits()
			if got != tt.want {
				t.Fatalf("getDigits() want %s, got %s", tt.want, got)
			}
		})
	}

}

func TestIsDisit(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  bool
	}{
		{name: "整数値", input: '1', want: true},
		{name: "アルファベット", input: 'a', want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isDigit(tt.input)
			if got != tt.want {
				t.Fatalf("isDigit() want %t, got %t", tt.want, got)
			}
		})
	}
}
