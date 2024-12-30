package lexer

import (
	"interpreter/pkg/token"
	"testing"
)

func TestNextToken(test *testing.T) {
	/*
		test1_input := `(+=(){},;)`

		tests  := []struct {
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
			{token.EOF, ""},
		}
	*/
	test2_input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`
	tests2 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.ASSIGN, "="},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	TEST_INPUT := test2_input
	TEST_COLLECTION := tests2
	_lexer := NewLexer(TEST_INPUT)

	for index, testToken := range TEST_COLLECTION {
		nextToken := _lexer.NextToken()
		if nextToken.Type != testToken.expectedType {
			test.Fatalf("tests[%d] - wrong tokentype. expected=%q, got=%q", index, testToken.expectedType, nextToken.Type)
		}
		if nextToken.Literal != testToken.expectedLiteral {
			test.Fatalf("tests[%d] - wrong tokenLiteral. expected=%q, got=%q", index, testToken.expectedLiteral, nextToken.Literal)
		}
	}
}
