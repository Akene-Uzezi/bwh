# BWH - Badwhipp Programming Language

A simple lexer/tokenizer, parser, and REPL implementation written in Go for the Badwhipp programming language.

## Features

- Tokenizes identifiers, integers, keywords (let, fn, true, false, if, else, return), operators (+, -, *, /, =, ==, !=, <, >, !), and punctuation
- Handles language constructs: variable declarations, function definitions, function calls, conditionals, boolean expressions
- Parser with AST generation for let statements and identifiers
- Includes comprehensive test coverage for lexer
- Interactive REPL for experimentation

## Project Structure

```
bwh/
├── go.mod              # Go module definition
├── main.go             # Entry point with REPL
├── commit.sh           # Commit helper script
├── lexer/
│   ├── lexer.go        # Lexer implementation
│   └── lexer_test.go   # Lexer tests
├── repl/
│   └── repl.go         # REPL implementation
├── token/
│   └── token.go        # Token types and structures
├── parser/
│   └── parser.go       # Parser implementation
└── ast/
    └── ast.go          # Abstract Syntax Tree nodes
```

## Token Types

- `ILLEGAL` - Unknown/invalid characters
- `EOF` - End of file
- `IDENT` - Identifiers (variables, functions)
- `INT` - Integer literals
- `ASSIGN` - `=`
- `PLUS` - `+`
- `MINUS` - `-`
- `BANG` - `!`
- `ASTERISK` - `*`
- `SLASH` - `/`
- `LT` - `<`
- `GT` - `>`
- `EQ` - `==`
- `NOT_EQ` - `!=`
- `COMMA` - `,`
- `SEMICOLON` - `;`
- `LPAREN` - `(`
- `RPAREN` - `)`
- `LBRACE` - `{`
- `RBRACE` - `}`
- `FUNCTION` - `fn` keyword
- `LET` - `let` keyword
- `TRUE` - `true` keyword
- `FALSE` - `false` keyword
- `IF` - `if` keyword
- `ELSE` - `else` keyword
- `RETURN` - `return` keyword

## Usage

### Running the REPL

```bash
go run main.go
```

This starts an interactive REPL where you can type Badwhipp code and see the tokenized output:

```
Hello user! This is the badwhipp programming language!
Feel free to type in commands
>>
```

### Using the Lexer Programmatically

```go
package main

import (
    "fmt"
    "bwh/lexer"
    "bwh/token"
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
        }
    }
}
```

### Using the Parser

```go
package main

import (
    "fmt"
    "bwh/lexer"
    "bwh/parser"
)

func main() {
    input := `let x = 5;`
    l := lexer.New(input)
    p := parser.New(l)
    
    program := p.ParseProgram()
    fmt.Printf("%+v\n", program)
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
!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

10 == 10;
10 != 9;
```

## Implementation Details

### Lexer (`lexer/lexer.go`)

The `Lexer` struct maintains the input string and current position. Key methods:

- `New(input string) *Lexer` - Creates a new lexer instance
- `NextToken() token.Token` - Returns the next token from input
- `readChar()` - Advances the position in the input
- `peekChar()` - Peeks at the next character without advancing
- `readIdentifier()` - Reads an identifier (variable/function name)
- `readNumber()` - Reads an integer literal
- `skipWhiteSpace()` - Skips whitespace characters

### Token Package (`token/token.go`)

Defines the token types and a lookup function for keywords:

- `TokenType` - String type for token classification
- `Token` - Struct containing Type and Literal
- `LookUpIdent(ident string) TokenType` - Maps keywords to token types

### Parser (`parser/parser.go`)

The `Parser` struct uses the lexer to parse tokens into an AST. Key methods:

- `New(l *lexer.Lexer) *Parser` - Creates a new parser instance
- `ParseProgram() *ast.Program` - Parses the entire program (currently returns nil, work in progress)
- `nextToken()` - Advances to the next token

### AST (`ast/ast.go`)

Defines the Abstract Syntax Tree node interfaces and implementations:

- `Node` interface - Base interface with `TokenLiteral()` method
- `Statement` interface - For statements like let declarations
- `Expression` interface - For expressions
- `Program` - Root node containing a list of statements
- `LetStatement` - Represents variable declarations (`let x = 5;`)
- `Identifier` - Represents variable/function names

### REPL (`repl/repl.go`)

Simple Read-Eval-Print Loop that:
- Reads input line by line
- Tokenizes each line using the lexer
- Prints each token

## Requirements

- Go 1.21+