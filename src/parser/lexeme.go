package parser

import (
	"./ast"
)

type LexemeSequence []Lexeme

type Lexeme struct {
	Token  ast.TokenKind
	Offset ast.TokenOffset
	Value  string
}
