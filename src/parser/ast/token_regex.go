package ast

import "regexp"

// TokenRegex represents a token with a value that is a regular expression.
type TokenRegex struct {
	Token
	Expression *regexp.Regexp
}

// Match will attempt to match the regular expression against the given input string.
// Must like TokenLiteral the value must be found at the beginning of the input string to be considered valid.
// Also note that regular expression matching can be costly at compute time, use literals where possible.
func (t TokenRegex) Match(input string) (bool, TokenOffset) {
	indexes := t.Expression.FindStringIndex(input)

	if len(indexes) == 0 {
		return false, NoTokenOffset
	}

	index := indexes[0]
	length := indexes[1] - index
	offset := TokenOffset{index, length}

	return true, offset
}
