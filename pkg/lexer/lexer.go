package lexer

import (
	"interpreter/pkg/token"
)

type Lexer struct {
	input   string
	currPos int
	nextPos int // current reading position in input (after current char)
	char    byte
}

func (lexer *Lexer) readCh() {
	if lexer.nextPos >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.nextPos]
	}
	lexer.currPos = lexer.nextPos
	lexer.nextPos += 1
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCh()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var _token token.Token

	switch lexer.char {
	case '=':
		_token = token.NewToken(token.ASSIGN, lexer.char)
	case ';':
		_token = token.NewToken(token.SEMICOLON, lexer.char)
	case ',':
		_token = token.NewToken(token.COMMA, lexer.char)
	case '(':
		_token = token.NewToken(token.LPAREN, lexer.char)
	case ')':
		_token = token.NewToken(token.RPAREN, lexer.char)
	case '{':
		_token = token.NewToken(token.LBRACE, lexer.char)
	case '}':
		_token = token.NewToken(token.RBRACE, lexer.char)
	case '+':
		_token = token.NewToken(token.PLUS, lexer.char)
	case 0:
		_token.Type = token.EOF
		_token.Literal = ""
	}

	lexer.readCh()
	return _token
}
