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
	if l.position != 0 {
		t.Fatalf("l.input want 0, got %d", l.position)
	}
}

func TestNextToken(t *testing.T) {
	l := New("1")
	token := l.NextToken()
	if token.Type != INT {
		t.Fatalf("l.NextToken().Literal want %d, got %d", INT, token.Type)
	}
	if token.Literal != "1" {
		t.Fatalf("l.NextToken().Literal want %s, got %s", "1", token.Literal)
	}
}
