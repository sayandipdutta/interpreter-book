module github.com/sayandipdutta/interpreter-book/parser

go 1.21.4

replace github.com/sayandipdutta/interpreter-book/token => ../token

replace github.com/sayandipdutta/interpreter-book/lexer => ../lexer

replace github.com/sayandipdutta/interpreter-book/ast => ../ast

require (
	github.com/sayandipdutta/interpreter-book/ast v0.0.0-00010101000000-000000000000
	github.com/sayandipdutta/interpreter-book/lexer v0.0.0-00010101000000-000000000000
	github.com/sayandipdutta/interpreter-book/token v0.0.0-00010101000000-000000000000
)
