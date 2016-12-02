package lexer

// checker receives a pointer to a lexer and tries to
// lex a token. If it is successful, it will return it as
// well as true, otherwise it will return an empty token and false
type checker interface {
	Check(*Lexer) (Token, bool)
}

type delimiterChecker struct{}

func (d delimiterChecker) Check(l *Lexer) (Token, bool) {
	c := l.ch

	switch l.ch {
	case ',':
		l.readChar()
		return newToken(COMMA, c, l.line), true
	case ';':
		l.readChar()
		return newToken(SEMICOLON, c, l.line), true
	case '(':
		l.readChar()
		return newToken(LPAREN, c, l.line), true
	case ')':
		l.readChar()
		return newToken(RPAREN, c, l.line), true
	case '{':
		l.readChar()
		return newToken(LBRACE, c, l.line), true
	case '}':
		l.readChar()
		return newToken(RBRACE, c, l.line), true
	}

	return Token{}, false
}

type eofChecker struct{}

func (c eofChecker) Check(l *Lexer) (Token, bool) {
	if l.ch == 0 {
		tok := Token{}
		tok.Type = EOF
		tok.Literal = ""
		l.readChar()
		return tok, true
	}

	return Token{}, false
}

type plusChecker struct{}

func (c plusChecker) Check(l *Lexer) (Token, bool) {
	if l.ch == '+' {
		l.readChar()
		if l.ch == '+' {
			tok := Token{}
			tok.Type = INC
			tok.Literal = "++"
			l.readChar()
			return tok, true
		}
		tok := newToken(PLUS, '+', l.line)
		return tok, true
	}

	return Token{}, false
}

type equalsChecker struct{}

func (c equalsChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if l.ch == '=' {
		if l.input[l.position:l.position+3] == "===" {
			tok.Type = IDENTITY
			tok.Literal = "==="
			l.advance(3)

			return tok, true
		}
		if l.input[l.position:l.position+2] == "==" {
			tok.Type = EQUALS
			tok.Literal = "=="
			l.advance(2)

			return tok, true
		}
		if l.input[l.position:l.position+2] == "=>" {
			tok.Type = ARROW
			tok.Literal = "=>"
			l.advance(2)

			return tok, true
		}
		if l.input[l.position:l.position+2] == "=&" {
			tok.Type = REFERENCE
			tok.Literal = "=&"
			l.advance(2)
			return tok, true
		}
		tok = newToken(ASSIGN, l.ch, l.line)
		l.readChar()
		return tok, true
	}

	return tok, false
}

type numberChecker struct{}

func (c numberChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if c.isInt(l.ch) {
		tok.Type = INT
		tok.Literal = c.readInteger(l)
		return tok, true
	}

	return tok, false
}

func (c numberChecker) isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func (c numberChecker) readInteger(l *Lexer) string {
	position := l.position
	for c.isInt(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

type identifierChecker struct{}

func (i identifierChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if i.isLetter(l.ch) {
		tok.Literal = i.readIdentifier(l)
		tok.Type = LookupIdent(tok.Literal)
		return tok, true
	}

	return tok, false
}

func (i identifierChecker) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '$'
}

func (i identifierChecker) readIdentifier(l *Lexer) string {
	position := l.position
	for i.isLetter(l.ch) || l.ch == '\\' {
		l.readChar()
	}

	return l.input[position:l.position]
}

type phptagChecker struct{}

func (p phptagChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if l.input[l.position:l.position+5] == "<?php" {
		tok.Type = PHPTAG
		tok.Literal = "<?php"
		for range tok.Literal {
			l.readChar()
		}
		return tok, true
	}

	return tok, false
}

type compareChecker struct{}

func (c compareChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if l.ch == '<' {
		if (l.input[l.position+1]) == '=' {
			tok.Type = LESSTHANOREQUAL
			tok.Literal = "<="
			l.advance(2)
			return tok, true
		}
		tok = newToken(LESSTHAN, l.ch, l.line)
		l.readChar()
		return tok, true
	}
	if l.ch == '>' {
		if (l.input[l.position+1]) == '=' {
			tok.Type = GREATERTHANOREQUAL
			tok.Literal = ">="
			l.advance(2)
			return tok, true
		}
		tok = newToken(GREATERTHAN, l.ch, l.line)
		l.readChar()
		return tok, true
	}

	return tok, false
}

type stringChecker struct {
	delimiter byte
	tokenType TokenType
}

func (s stringChecker) Check(l *Lexer) (Token, bool) {
	tok := Token{}
	if l.ch != s.delimiter {
		return tok, false
	}
	l.readChar()
	tok.Type = s.tokenType
	tok.Literal = s.readString(l)
	l.readChar()

	return tok, true
}

func (s stringChecker) readString(l *Lexer) string {
	pos := l.position
	var res []byte
	for {
		l.scan([]byte{s.delimiter, '\\'})
		if l.ch == '\\' {
			res = append(res, l.input[pos:l.position]...)
			l.advance(2)
			pos = l.position - 1
			continue
		}
		res = append(res, l.input[pos:l.position]...)
		break
	}
	return string(res)
}
