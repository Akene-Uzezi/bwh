package main

import (
	"bwh/lexer"
	"bwh/token"
	"fmt"
)

func main() {
	input := `
		let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}
	`
	l := lexer.New(input)
	for {
		l.NextToken()
		fmt.Printf("TokenType: %s, TokenLiteral: %s\n", l.NextToken().Type, l.NextToken().Literal)
		if l.NextToken().Type == token.EOF {
			break
		} else {
			continue
		}
	}
}
