package repl

import (
	"bufio"
	"fmt"
	"github.com/bestform/shmehashme/lexer"
	"io"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
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
			fmt.Printf("%+v\n", tok)
		}
	}
}
