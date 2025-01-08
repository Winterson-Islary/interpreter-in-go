package parser

import (
	"interpreter/pkg/ast"
	"interpreter/pkg/lexer"
	"testing"
)

func TestLetStatements(test *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 90536;
	`
	_lexer := lexer.NewLexer(input)
	parser := NewParser(_lexer)
	program := parser.ParseProgram()

	if program == nil {
		test.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		test.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, item := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(test, stmt, item.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(test *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		test.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		test.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		test.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		test.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
