package parser

import (
	"testing"

	"github.com/sayandipdutta/interpreter-book/ast"
	"github.com/sayandipdutta/interpreter-book/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got: [%d]", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, tt.expectedIdentifier, stmt) {
			return
		}
	}
}

func testLetStatement(t *testing.T, expected string, got ast.Statement) bool {
	if got.TokenLiteral() != "let" {
		t.Errorf("Expected 'let', got [%q] as TokenLiteral", got.TokenLiteral())
	}
	letStmt, ok := got.(*ast.LetStatement)
	if !ok {
		t.Errorf("not a *ast.LetStatement, got=%T", got)
		return false
	}

	if letStmt.Name.Value != expected {
		t.Errorf("Expected letStmt.Name.Value='%s', got='%s'", expected, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expected {
		t.Errorf("Expected letStmt.Name.TokenLiteral='%s', got='%s'", expected, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
