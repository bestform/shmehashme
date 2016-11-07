package lexer

import (
	"testing"

	"os"

	"strconv"

	"github.com/bestform/shmehashme/token"
)

type testcase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

type testsetup struct {
	filename  string
	testcases []testcase
}

var tests = []testsetup{
	{
		filename: "fixtures/loopsAndConditions.php",
		testcases: []testcase{
			{token.PHPTAG, "<?php"},

			// if
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.TRUE, "true"},
			{token.IDENTITY, "==="},
			{token.FALSE, "false"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "$baz"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},

			// for
			{token.FOR, "for"},
			{token.LPAREN, "("},
			{token.IDENT, "$i"},
			{token.ASSIGN, "="},
			{token.INT, "0"},
			{token.SEMICOLON, ";"},
			{token.IDENT, "$i"},
			{token.LESSTHAN, "<"},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.IDENT, "$i"},
			{token.INC, "++"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
		},
	},
	{
		filename: "fixtures/classStructure.php",
		testcases: []testcase{
			{token.PHPTAG, "<?php"},

			{token.USE, "use"},
			{token.IDENT, "Foo\\Bar"},
			{token.SEMICOLON, ";"},

			{token.CLASS, "class"},
			{token.IDENT, "Foo"},
			{token.LBRACE, "{"},

			{token.PUBLIC, "public"},
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

			{token.PRIVATE, "private"},
			{token.FUNCTION, "function"},
			{token.IDENT, "bar"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},

			{token.PROTECTED, "protected"},
			{token.FUNCTION, "function"},
			{token.IDENT, "baz"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},

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
		},
	},
}

func TestNextToken(t *testing.T) {

	for _, testcase := range tests {
		input, err := os.OpenFile(testcase.filename, os.O_RDONLY, 0666)
		defer input.Close()
		if err != nil {
			t.Fatal("error opening test fixture")
		}
		l, err := New(input)
		if err != nil {
			t.Fatal("error creating lexer", err)
		}

		for i, tt := range testcase.testcases {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q (%v)",
					i, tt.expectedType, tok.Type, tok.Literal)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
					i, tt.expectedLiteral, tok.Literal)
			}
		}
	}

}

func TestLineSupport(t *testing.T) {

	input, err := os.OpenFile("fixtures/lineNumbers.php", os.O_RDONLY, 0666)
	defer input.Close()
	if err != nil {
		t.Fatal("Error reading fixture", err)
	}

	l, err := New(input)
	if err != nil {
		t.Fatal("error creating lexer", err)
	}

	var tok token.Token
	for tok.Type != token.EOF {
		tok = l.NextToken()
		if tok.Type == token.INT {
			expectedLine, err := strconv.Atoi(tok.Literal)
			if err != nil {
				t.Fatal("error reading integer token")
			}
			if tok.Line != expectedLine {
				t.Fatalf("tests %v, expected line to be %v but got %v", tok, expectedLine, tok.Line)
			}
		}
	}

}
