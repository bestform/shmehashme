package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/bestform/shmehashme/lexer"
)

const prompt = ">> "

// Start will launch a simple REPL that takes source code and will display the lexer result
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		reader := strings.NewReader(line)

		l, err := lexer.New(reader)
		if err != nil {
			return
		}

		for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
			fmt.Println(tok)
		}
	}
}
