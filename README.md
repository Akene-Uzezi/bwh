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
    "bwh/token"
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

## Implementation Details

### Lexer (`lexer/lexer.go`)

The `Lexer` struct maintains the input string and current position. Key methods:

- `New(input string) *Lexer` - Creates a new lexer instance
- `NextToken() token.Token` - Returns the next token from input
- `readChar()` - Advances the position in the input
- `readIdentifier()` - Reads an identifier (variable/function name)
- `readNumber()` - Reads an integer literal
- `skipWhiteSpace()` - Skips whitespace characters

### Token Package (`token/token.go`)

Defines the token types and a lookup function for keywords:

- `TokenType` - String type for token classification
- `Token` - Struct containing Type and Literal
- `LookUpIdent(ident string) TokenType` - Maps keywords to token types

## Requirements

- Go 1.21+
