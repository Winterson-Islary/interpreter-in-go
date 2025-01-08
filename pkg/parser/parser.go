package parser

import (
	"interpreter/pkg/ast"
	"interpreter/pkg/lexer"
	"interpreter/pkg/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.currToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}
func (parser *Parser) ParseProgram() *ast.Program {
	return nil
}
