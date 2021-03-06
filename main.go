package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bestform/shmehashme/lexer"
	"github.com/bestform/shmehashme/repl"
)

func main() {
	fmt.Println("Why hello there! This a REPL for the shmehashme PHP lexer")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

func mainFile() {
	file := flag.String("file", "", "File to lex")
	flag.Parse()

	if *file == "" {
		fmt.Println("Please specify a file")
		os.Exit(1)
	}

	f, err := os.Open(*file)
	defer f.Close()
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	l, err := lexer.New(f)
	if err != nil {
		fmt.Println("Error creating lexer: ", err)
		os.Exit(1)
	}
	var tokens []lexer.Token
	var t lexer.Token
	for {
		t = l.NextToken()
		if t.Type == lexer.EOF {
			break
		}
		tokens = append(tokens, t)
	}

	o := lexer.PrettyPrint(tokens)
	fmt.Print(o)
}
