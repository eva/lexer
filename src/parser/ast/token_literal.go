package ast

import "strings"

type TokenLiteral struct {
	Token
	Literal string
}

func (t TokenLiteral) Match(input string) (bool, TokenOffset) {
	index := strings.Index(input, t.Literal)

	if index == -1 {
		return false, NoTokenOffset
	}

	length := len(t.Literal)
	offset := TokenOffset{index, length}

	return true, offset
}
