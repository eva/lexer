package parser

import (
	"./ast"
)

type lexeme struct {
	token  ast.Token
	offset ast.TokenOffset
	value  string
}

func NewLexeme(token ast.Token, offset ast.TokenOffset, value string) Lexeme {
	lexeme := lexeme{}
	lexeme.token = token
	lexeme.offset = offset
	lexeme.value = value

	return lexeme
}

func (lexeme lexeme) Token() ast.Token {
	return lexeme.token
}

func (lexeme lexeme) Offset() ast.TokenOffset {
	return lexeme.offset
}

func (lexeme lexeme) Value() string {
	return lexeme.value
}
