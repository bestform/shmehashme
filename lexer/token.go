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
	// ILLEGAL represents a not character the lexer does not know
	ILLEGAL = "ILLEGAL"
	// EOF indicates the end of the file
	EOF = "EOF"

	// Identifiers and literals

	// IDENT represents all kinds of identifiers, like variables, function names, internal functions etc.
	IDENT = "IDENT" // $foo, $bar, foo, print
	// INT is an integer
	INT = "INT" // 123456

	// Operators

	// ASSIGN is "="
	ASSIGN = "="
	// PLUS is "+""
	PLUS = "+"
	// EQUALS is "=="
	EQUALS = "=="
	// IDENTITY is "==="
	IDENTITY = "==="
	// LESSTHAN is "<"
	LESSTHAN = "<"
	// INC is "++" - used for both preinc and postinc. It is the job of the parser to differentiate
	INC = "++"

	// Delimiters

	// COMMA is ","
	COMMA = ","
	// SEMICOLON is ";"
	SEMICOLON = ";"
	// LPAREN is "("
	LPAREN = "("
	// RPAREN is ")"
	RPAREN = ")"
	// LBRACE is "{"
	LBRACE = "{"
	// RBRACE is "}"
	RBRACE = "}"

	// Keywords

	// FUNCTION is "function"
	FUNCTION = "function"
	// RETURN is "return"
	RETURN = "return"
	// PUBLIC is "public"
	PUBLIC = "public"
	// PRIVATE is "private"
	PRIVATE = "private"
	// PROTECTED is "protected"
	PROTECTED = "protected"
	// PHPTAG is "<?php"
	PHPTAG = "<?php"
	// CLASS is "class"
	CLASS = "class"
	// IF is "if"
	IF = "if"
	// TRUE is "true"
	TRUE = "true"
	// FALSE is "false"
	FALSE = "false"
	// USE is "use"
	USE = "use"
	// FOR is "for"
	FOR = "for"
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
