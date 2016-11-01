package lexer

import (
	"bufio"
	"io"

	"github.com/bestform/shmehashme/token"
)

// Lexer can lex PHP source code into tokens
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New will return a pointer to a fresh lexer initialized with input
func New(input io.Reader) (*Lexer, error) {

	reader := bufio.NewReader(input)
	var delim byte
	stringInput, err := reader.ReadString(delim)
	if err != nil && err != io.EOF {
		return nil, err
	}

	l := &Lexer{input: stringInput}
	l.readChar()

	return l, nil
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) advance(p int) {
	for i := 0; i < p; i++ {
		l.readChar()
	}
}

// NextToken returns the next token and advances internally. At the end it will return token.EOF
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.input[l.position:l.position+3] == "===" {
			tok.Type = token.IDENTITY
			tok.Literal = "==="
			l.advance(3)

			return tok
		}
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isInt(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readInteger()
			return tok
		}

		if l.input[l.position:l.position+5] == "<?php" {
			tok.Type = token.PHPTAG
			tok.Literal = "<?php"
			for _ = range tok.Literal {
				l.readChar()
			}
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)

	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
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

func newToken(t token.TokenType, l byte) token.Token {
	return token.Token{Type: t, Literal: string(l)}
}
