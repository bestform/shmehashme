package lexer

import (
	"bufio"
	"io"
)

// Lexer can lex PHP source code into tokens
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line in input
}

// New will return a pointer to a fresh lexer initialized with input
func New(input io.Reader) (*Lexer, error) {

	reader := bufio.NewReader(input)
	var delim byte
	stringInput, err := reader.ReadString(delim)
	if err != nil && err != io.EOF {
		return nil, err
	}

	l := &Lexer{input: stringInput, line: 1}
	l.readChar()

	return l, nil
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	if l.ch == '\n' {
		l.line++
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) advance(p int) {
	for i := 0; i < p; i++ {
		l.readChar()
	}
}

// NextToken returns the next token and advances internally. At the end it will return EOF
func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()
	tok.Line = l.line

	delimiterToken, ok := handleDelimiters(l)
	if ok {
		delimiterToken.Line = l.line
		return delimiterToken
	}

	if l.ch == 0 {
		tok.Type = EOF
		tok.Literal = ""
		l.readChar()
		return tok
	}

	if l.ch == '+' {
		l.readChar()
		if l.ch == '+' {
			tok.Type = INC
			tok.Literal = "++"
			l.readChar()
			return tok
		}
		tok = newToken(PLUS, '+', l.line)
		return tok
	}

	if l.ch == '=' {
		if l.input[l.position:l.position+3] == "===" {
			tok.Type = IDENTITY
			tok.Literal = "==="
			l.advance(3)

			return tok
		}
		if l.input[l.position:l.position+2] == "==" {
			tok.Type = EQUALS
			tok.Literal = "=="
			l.advance(2)

			return tok
		}
		if l.input[l.position:l.position+2] == "=>" {
			tok.Type = ARROW
			tok.Literal = "=>"
			l.advance(2)

			return tok
		}
		tok = newToken(ASSIGN, l.ch, l.line)
		l.readChar()
		return tok
	}

	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = LookupIdent(tok.Literal)
		return tok
	}

	if isInt(l.ch) {
		tok.Type = INT
		tok.Literal = l.readInteger()
		return tok
	}

	if l.input[l.position:l.position+5] == "<?php" {
		tok.Type = PHPTAG
		tok.Literal = "<?php"
		for range tok.Literal {
			l.readChar()
		}
		return tok
	}

	if l.ch == '<' {
		tok = newToken(LESSTHAN, l.ch, l.line)
		l.readChar()
		return tok
	}

	tok = newToken(ILLEGAL, l.ch, l.line)
	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isBackslash(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readInteger() string {
	position := l.position
	for isInt(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func isInt(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isBackslash(ch byte) bool {
	return ch == '\\'
}

func newToken(t TokenType, l byte, line int) Token {
	return Token{Type: t, Literal: string(l), Line: line}
}

func handleDelimiters(l *Lexer) (Token, bool) {
	c := l.ch
	if l.ch == ',' {
		l.readChar()
		return newToken(COMMA, c, l.line), true
	} else if l.ch == ';' {
		l.readChar()
		return newToken(SEMICOLON, c, l.line), true
	} else if l.ch == '(' {
		l.readChar()
		return newToken(LPAREN, c, l.line), true
	} else if l.ch == ')' {
		l.readChar()
		return newToken(RPAREN, c, l.line), true
	} else if l.ch == '{' {
		l.readChar()
		return newToken(LBRACE, c, l.line), true
	} else if l.ch == '}' {
		l.readChar()
		return newToken(RBRACE, c, l.line), true
	}

	return Token{}, false
}
