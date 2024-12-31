package token

//? Defining structure of a Token and its types

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

func NewToken(tokenType TokenType, literal byte) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}

var KEYWORDS = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifierType(identifier string) TokenType {
	if _tokenType, ok := KEYWORDS[identifier]; ok {
		return _tokenType
	}
	return IDENT
}
