package lexer

type Lexer struct {
	input    string
	currPos  int
	nextPos  int // current reading position in input (after current char)
	currChar byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}
