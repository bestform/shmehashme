package lexer

import (
	"os"
	"strconv"
	"testing"
)

type testcase struct {
	expectedType    TokenType
	expectedLiteral string
}

type testsetup struct {
	filename  string
	testcases []testcase
}

var tests = []testsetup{
	{
		filename: "fixtures/strings.php",
		testcases: []testcase{
			{PHPTAG, "<?php"},

			{DOUBLEQUOTEDSTRING, "foo"},
			{SEMICOLON, ";"},
			{DOUBLEQUOTEDSTRING, "foo\"bar"},
			{SEMICOLON, ";"},

			{SINGLEQUOTEDSTRING, "foo"},
			{SEMICOLON, ";"},
			{SINGLEQUOTEDSTRING, "foo'bar"},
			{SEMICOLON, ";"},
		},
	},
	{
		filename: "fixtures/assignments.php",
		testcases: []testcase{
			{PHPTAG, "<?php"},

			{IDENT, "$foo"},
			{ASSIGN, "="},
			{IDENT, "$bar"},
			{SEMICOLON, ";"},

			{IDENT, "$foo"},
			{REFERENCE, "=&"},
			{IDENT, "$bar"},
			{SEMICOLON, ";"},
		},
	},
	{
		filename: "fixtures/loopsAndConditions.php",
		testcases: []testcase{
			{PHPTAG, "<?php"},

			// if with equals check
			{IF, "if"},
			{LPAREN, "("},
			{INT, "1337"},
			{EQUALS, "=="},
			{INT, "42"},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			// if with identity check
			{IF, "if"},
			{LPAREN, "("},
			{TRUE, "true"},
			{IDENTITY, "==="},
			{FALSE, "false"},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			// for
			{FOR, "for"},
			{LPAREN, "("},
			{IDENT, "$i"},
			{ASSIGN, "="},
			{INT, "0"},
			{SEMICOLON, ";"},
			{IDENT, "$i"},
			{LESSTHAN, "<"},
			{INT, "10"},
			{SEMICOLON, ";"},
			{IDENT, "$i"},
			{INC, "++"},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			// foreach
			{FOREACH, "foreach"},
			{LPAREN, "("},
			{IDENT, "$i"},
			{AS, "as"},
			{IDENT, "$j"},
			{ARROW, "=>"},
			{IDENT, "$k"},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			// compare
			{IDENT, "$i"},
			{LESSTHAN, "<"},
			{INT, "10"},
			{SEMICOLON, ";"},

			{IDENT, "$i"},
			{GREATERTHAN, ">"},
			{INT, "10"},
			{SEMICOLON, ";"},

			{IDENT, "$i"},
			{LESSTHANOREQUAL, "<="},
			{INT, "10"},
			{SEMICOLON, ";"},

			{IDENT, "$i"},
			{GREATERTHANOREQUAL, ">="},
			{INT, "10"},
			{SEMICOLON, ";"},
		},
	},
	{
		filename: "fixtures/classStructure.php",
		testcases: []testcase{
			{PHPTAG, "<?php"},

			{USE, "use"},
			{IDENT, "Foo\\Bar"},
			{SEMICOLON, ";"},

			{CLASS, "class"},
			{IDENT, "Foo"},
			{EXTENDS, "extends"},
			{IDENT, "Bar"},
			{IMPLEMENTS, "implements"},
			{IDENT, "BarInterface"},
			{LBRACE, "{"},

			{PUBLIC, "public"},
			{FUNCTION, "function"},
			{IDENT, "foo"},
			{LPAREN, "("},
			{IDENT, "$bar"},
			{COMMA, ","},
			{IDENT, "$baz"},
			{RPAREN, ")"},

			{LBRACE, "{"},

			{RETURN, "return"},
			{IDENT, "$bar"},
			{PLUS, "+"},
			{IDENT, "$baz"},
			{SEMICOLON, ";"},
			{RBRACE, "}"},

			{STATIC, "static"},
			{PRIVATE, "private"},
			{FUNCTION, "function"},
			{IDENT, "bar"},
			{LPAREN, "("},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			{PROTECTED, "protected"},
			{FUNCTION, "function"},
			{IDENT, "baz"},
			{LPAREN, "("},
			{RPAREN, ")"},
			{LBRACE, "{"},
			{RBRACE, "}"},

			{RBRACE, "}"},

			{IDENT, "print"},
			{LPAREN, "("},
			{IDENT, "foo"},
			{LPAREN, "("},
			{INT, "1"},
			{COMMA, ","},
			{INT, "2"},
			{RPAREN, ")"},
			{RPAREN, ")"},
			{SEMICOLON, ";"},
			{EOF, ""},
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

	var tok Token
	for tok.Type != EOF {
		tok = l.NextToken()
		if tok.Type == INT {
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
