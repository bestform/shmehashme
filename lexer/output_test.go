package lexer

import (
	"testing"
)

func TestPrettyPrint(t *testing.T) {

	input := []Token{
		{PHPTAG, "<?php", 1},
		{IDENT, "$foo", 2},
		{ASSIGN, "=", 2},
		{IDENT, "$bar", 2},
		{SEMICOLON, ";", 2},
	}

	expectedOutputWithLines := "1\tPHPTAG\t<?php\n" +
		"2\tIDENT\t$foo\n" +
		"2\tASSIGN\t=\n" +
		"2\tIDENT\t$bar\n" +
		"2\tSEMICOLON\t;\n"

	outputWithLines := PrettyPrint(input)

	if outputWithLines != expectedOutputWithLines {
		t.Fatalf("Output (with lines) does not match expected output:\nEXPECTED:\n%s\n\nACTUAL:\n%s", expectedOutputWithLines, outputWithLines)
	}

}
