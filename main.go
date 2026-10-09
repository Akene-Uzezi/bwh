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

	10 == 10;
	10 != 9;
	`
	l := lexer.New(input)
	for {
		tok := l.NextToken()
		fmt.Printf("TokenType: %s, TokenLiteral: %s\n", tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		} else {
			continue
		}
	}
}
