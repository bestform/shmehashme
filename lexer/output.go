package lexer

import (
	"bytes"
	"fmt"
)

// PrettyPrint prints the list of token in a sane way
// (can be used for example for output on the command line)
func PrettyPrint(t []Token) string {
	var b bytes.Buffer

	for _, tok := range t {
		b.WriteString(fmt.Sprintf("%v\t%s\t%s\n", tok.Line, tok.Type, tok.Literal))
	}

	return b.String()
}
