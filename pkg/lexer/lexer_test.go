package lexer

import (
	"interpreter/pkg/token"
	"testing"
)

func TestNextToken(test *testing.T) {
	input := `(+=(){},;)`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.EOF, "EOF"},
	}

	_lexer := NewLexer(input)

	for index, testToken := range tests {
		nextToken := _lexer.NextToken()
		if nextToken.Type != testToken.expectedType {
			test.Fatalf("tests[%d] - wrong tokentype. expected=%q, got=%q", index, testToken.expectedType, nextToken.Type)
		}
		if nextToken.Literal != testToken.expectedLiteral {
			test.Fatalf("tests[%d] - wrong tokenLiteral. expected=%q, got=%q", index, testToken.expectedLiteral, nextToken.Literal)
		}
	}
}
