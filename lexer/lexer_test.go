package lexer

import (
	"shmehashme/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	function foo($bar, $baz)
	{
    	return $bar + $baz;
	}

	print(foo(1, 2));
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
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

	l := New(input)

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
