package lexer

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// Lexer can lex PHP source code into tokens
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current rune under examination
	chsize       int  // current length of the rune in ch
	line         int  // current line in input
	checkers     []checker
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
	l.checkers = []checker{
		delimiterChecker{},
		eofChecker{},
		plusChecker{},
		equalsChecker{},
		numberChecker{},
		identifierChecker{},
		phptagChecker{},
		compareChecker{},
		stringChecker{
			delimiter: '"',
			tokenType: DOUBLEQUOTEDSTRING,
		},
		stringChecker{
			delimiter: '\'',
			tokenType: SINGLEQUOTEDSTRING,
		},
	}
	l.readChar()

	return l, nil
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch, l.chsize = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	if l.ch == '\n' {
		l.line++
	}
	l.position = l.readPosition
	l.readPosition += l.chsize
}

func (l *Lexer) advance(p int) {
	for i := 0; i < p; i++ {
		l.readChar()
	}
}

func (l *Lexer) scan(c []rune) {
	for !runeInSlice(l.ch, c) {
		l.readChar()
		if l.ch == 0 {
			return
		}
	}
}

func runeInSlice(b rune, s []rune) bool {
	for _, test := range s {
		if b == test {
			return true
		}
	}

	return false
}

// NextToken returns the next token and advances internally. At the end it will return EOF
func (l *Lexer) NextToken() Token {

	l.skipWhitespace()

	for _, c := range l.checkers {
		if tok, ok := c.Check(l); ok {
			tok.Line = l.line
			return tok
		}
	}

	var tok Token
	tok.Line = l.line
	tok = newToken(ILLEGAL, l.ch, l.line)
	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(t TokenType, l rune, line int) Token {
	return Token{Type: t, Literal: string(l), Line: line}
}
