# BWH - Basic Lexer/Tokenizer in Go

A simple lexer implementation written in Go for tokenizing a basic programming language.

## Features

- Tokenizes identifiers, integers, keywords (let, fn), operators (+, =), and punctuation
- Handles basic language constructs: variable declarations, function definitions, function calls
- Includes comprehensive test coverage

## Project Structure

```
bwh/
├── go.mod              # Go module definition
├── lexer/
│   ├── lexer.go        # Lexer implementation
│   └── lexer_test.go   # Lexer tests
└── token/
    └── token.go        # Token types and structures
```

## Token Types

- `ILLEGAL` - Unknown/invalid characters
- `EOF` - End of file
- `IDENT` - Identifiers (variables, functions)
- `INT` - Integer literals
- `ASSIGN` - `=`
- `PLUS` - `+`
- `COMMA` - `,`
- `SEMICOLON` - `;`
- `LPAREN` - `(`
- `RPAREN` - `)`
- `LBRACE` - `{`
- `RBRACE` - `}`
- `FUNCTION` - `fn` keyword
- `LET` - `let` keyword

## Usage

```go
package main

import (
    "fmt"
    "bwh/lexer"
)

func main() {
    input := `let x = 5;`
    l := lexer.New(input)
    
    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
        fmt.Printf("%s: %s\n", tok.Type, tok.Literal)
    }
}
```

## Running Tests

```bash
go test ./...
```

## Building

```bash
go build
```

## Example Input

The lexer can tokenize code like:

```javascript
let five = 5;
let ten = 10;
let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
```

## License

MIT