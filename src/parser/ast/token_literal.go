package ast

import "strings"

// TokenLiteral represents a literal string that should be matched in the sequence.
type TokenLiteral struct {
	Token
	Literal string
}

// Match will attempt to match the literal value against the given input.
// The literal might match but it must be found at the beginning of the input string to be valid.
func (t TokenLiteral) Match(input string) (bool, TokenOffset) {
	index := strings.Index(input, t.Literal)

	if index == -1 {
		return false, NoTokenOffset
	}

	length := len(t.Literal)
	offset := TokenOffset{index, length}

	return true, offset
}
