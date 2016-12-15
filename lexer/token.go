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
	// DOUBLEQUOTEDSTRING represents double quoted strings
	DOUBLEQUOTEDSTRING = "DOUBLEQUOTEDSTRING"
	// SINGLEQUOTEDSTRING represents double quoted strings
	SINGLEQUOTEDSTRING = "SINGLEQUOTEDSTRING"
	// COMMENT is a comment
	COMMENT = "COMMENT"

	// Identifiers and literals

	// IDENT represents all kinds of identifiers, like variables, function names, internal functions etc.
	IDENT = "IDENT" // $foo, $bar, foo, print
	// INT is an integer
	INT = "INT" // 123456

	// Operators

	// ASSIGN is "="
	ASSIGN = "ASSIGN"
	// REFERENCE is "=&"
	REFERENCE = "REFERENCE"
	// PLUS is "+"
	PLUS = "PLUS"
	// MINUS is "-"
	MINUS = "MINUS"
	// MULTIPLY is "*"
	MULTIPLY = "MULTIPLY"
	// DIVIDE is "/"
	DIVIDE = "DIVIDE"
	// EQUALS is "=="
	EQUALS = "EQUALS"
	// IDENTITY is "==="
	IDENTITY = "IDENTITY"
	// LESSTHAN is "<"
	LESSTHAN = "LESSTHAN"
	// GREATERTHAN is ">"
	GREATERTHAN = "GREATERTHAN"
	// LESSTHANOREQUAL is "<="
	LESSTHANOREQUAL = "LESSTHANOREQUAL"
	// GREATERTHANOREQUAL is ">="
	GREATERTHANOREQUAL = "GREATERTHANOREQUAL"
	// OR is "||"
	OR = "OR"
	// AND is "&&"
	AND = "AND"
	// NOT is "!"
	NOT = "NOT"
	// INC is "++" - used for both preinc and postinc. It is the job of the parser to differentiate
	INC = "INC"
	// DEC is "--" - used for both preinc and postinc. It is the job of the parser to differentiate
	DEC = "DEC"
	// QUESTIONMARK is "?" - as used in ternary operations
	QUESTIONMARK = "QUESTIONMARK"
	// COLON is ":" - as used in ternary operations
	COLON = "COLON"

	// Delimiters

	// COMMA is ","
	COMMA = "COMMA"
	// DOT is "."
	DOT = "DOT"
	// SEMICOLON is ";"
	SEMICOLON = "SEMICOLON"
	// LPAREN is "("
	LPAREN = "LPAREN"
	// RPAREN is ")"
	RPAREN = "RPAREN"
	// LBRACE is "{"
	LBRACE = "LBRACE"
	// RBRACE is "}"
	RBRACE = "RBRACE"
	// LSQUAREBRACKET is "["
	LSQUAREBRACKET = "LSQUAREBRACKET"
	// RSQUAREBRACKET is "]"
	RSQUAREBRACKET = "RSQUAREBRACKET"

	// Keywords

	// FUNCTION is "function"
	FUNCTION = "FUNCTION"
	// RETURN is "return"
	RETURN = "RETURN"
	// PUBLIC is "public"
	PUBLIC = "PUBLIC"
	// PRIVATE is "private"
	PRIVATE = "PRIVATE"
	// STATIC is "static"
	STATIC = "STATIC"
	// PROTECTED is "protected"
	PROTECTED = "PROTECTED"
	// PHPTAG is "<?php"
	PHPTAG = "PHPTAG"
	// CLASS is "class"
	CLASS = "CLASS"
	// IMPLEMENTS is "implements"
	IMPLEMENTS = "IMPLEMENTS"
	// EXTENDS is "extends"
	EXTENDS = "EXTENDS"
	// IF is "if"
	IF = "IF"
	// TRUE is "true"
	TRUE = "TRUE"
	// FALSE is "false"
	FALSE = "FALSE"
	// USE is "use"
	USE = "USE"
	// FOR is "for"
	FOR = "FOR"
	// FOREACH is "foreach"
	FOREACH = "FOREACH"
	// AS is as as used in a foreach loop
	AS = "AS"
	// DOUBLEARROW is => as used in a foreach loop
	DOUBLEARROW = "DOUBLEARROW"
	// ARROW is -> as used in attribute access
	ARROW = "ARROW"

	// PHP7

	// SPACESHIP is the spaceship operator added in php7: <=>
	SPACESHIP = "SPACESHIP"
)

var keywords = map[string]TokenType{
	"function":   FUNCTION,
	"return":     RETURN,
	"public":     PUBLIC,
	"private":    PRIVATE,
	"protected":  PROTECTED,
	"static":     STATIC,
	"<?php":      PHPTAG,
	"class":      CLASS,
	"if":         IF,
	"true":       TRUE,
	"false":      FALSE,
	"use":        USE,
	"for":        FOR,
	"foreach":    FOREACH,
	"as":         AS,
	"implements": IMPLEMENTS,
	"extends":    EXTENDS,
}

// LookupIdent will search for possible keywords and return the
// appropriate TokenType
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
