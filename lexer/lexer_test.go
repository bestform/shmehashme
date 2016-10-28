package lexer

import (
	"testing"

	"os"

	"github.com/bestform/shmehashme/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PUBLIC, "public"},
		{token.PRIVATE, "private"},
		{token.PROTECTED, "protected"},
		{token.FUNCTION, "function"},
		{token.IDENT, "foo"},
		{token.LPAREN, "("},
		{token.IDENT, "$bar"},
		{token.COMMA, ","},
		{token.IDENT, "$baz"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "$bar"},
		{token.PLUS, "+"},
		{token.IDENT, "$baz"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IDENT, "print"},
		{token.LPAREN, "("},
		{token.IDENT, "foo"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	input, err := os.OpenFile("fixtures/test.php", os.O_RDONLY, 0666)
	defer input.Close()

	if err != nil {
		t.Fatal("error opening test fixture")
	}
	l, err := New(input)
	if err != nil {
		t.Fatal("error creating lexer", err)
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
