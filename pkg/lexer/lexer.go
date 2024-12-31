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

func (lexer *Lexer) peekCh() byte {
	if lexer.nextPos >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.nextPos]
	}
}

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.currPos
	for isLetter(lexer.char) {
		lexer.readCh()
	}
	return lexer.input[startPosition:lexer.currPos]
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.currPos
	for isDigit(lexer.char) {
		lexer.readCh()
	}
	return lexer.input[startPosition:lexer.currPos]
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCh()
	return lexer
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readCh()
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var _token token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		if lexer.peekCh() == '=' {
			prevChar := lexer.char
			lexer.readCh()
			literal := string(prevChar) + string(lexer.char)
			_token = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			_token = token.NewToken(token.ASSIGN, lexer.char)
		}
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
	case '-':
		_token = token.NewToken(token.MINUS, lexer.char)
	case '!':
		if lexer.peekCh() == '=' {
			prevChar := lexer.char
			lexer.readCh()
			literal := string(prevChar) + string(lexer.char)
			_token = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			_token = token.NewToken(token.BANG, lexer.char)
		}
	case '*':
		_token = token.NewToken(token.ASTERISK, lexer.char)
	case '/':
		_token = token.NewToken(token.SLASH, lexer.char)
	case '<':
		_token = token.NewToken(token.LT, lexer.char)
	case '>':
		_token = token.NewToken(token.GT, lexer.char)
	case 0:
		_token.Type = token.EOF
		_token.Literal = ""
	default:
		if isLetter(lexer.char) {
			_token.Literal = lexer.readIdentifier()
			_token.Type = token.LookupIdentifierType(_token.Literal)
			return _token
		} else if isDigit(lexer.char) {
			_token.Literal = lexer.readNumber()
			_token.Type = token.INT
			return _token
		} else {
			_token = token.NewToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readCh()
	return _token
}
