package lexer

type Lexer struct {
	input        string
	position     int  // curent position in input
	readPosition int  // current readPosition in input
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
