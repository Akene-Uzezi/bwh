// Package repl
package repl

import (
	"bufio"
	"bwh/lexer"
	"bwh/token"
	"fmt"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
