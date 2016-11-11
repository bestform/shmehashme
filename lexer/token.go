package lexer

// TokenType defines the type of a Token (see below)
type TokenType string

// Token represents one token of a pecific type and its literal representation
type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // $foo, $bar, foo, print
	INT   = "INT"   // 123456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	EQUALS   = "=="
	IDENTITY = "==="
	LESSTHAN = "<"
	INC      = "++"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION  = "function"
	RETURN    = "return"
	PUBLIC    = "public"
	PRIVATE   = "private"
	PROTECTED = "protected"
	PHPTAG    = "<?php"
	CLASS     = "class"
	IF        = "if"
	TRUE      = "true"
	FALSE     = "false"
	USE       = "use"
	FOR       = "for"
)

var keywords = map[string]TokenType{
	"function":  FUNCTION,
	"return":    RETURN,
	"public":    PUBLIC,
	"private":   PRIVATE,
	"protected": PROTECTED,
	"<?php":     PHPTAG,
	"class":     CLASS,
	"if":        IF,
	"true":      TRUE,
	"false":     FALSE,
	"use":       USE,
	"for":       FOR,
}

// LookupIdent will search for possible keywords and return the
// appropriate TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
