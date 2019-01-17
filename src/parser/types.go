package parser

import "./ast"

type Lexeme interface {
	Token() ast.Token
	Offset() ast.TokenOffset
	Value() string
}

type LexemeSequence []Lexeme
