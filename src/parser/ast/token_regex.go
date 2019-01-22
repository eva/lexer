package ast

import "regexp"

type TokenRegex struct {
	Token
	Expression *regexp.Regexp
}

func (t TokenRegex) Match(input string) (bool, TokenOffset) {
	indexes := t.Expression.FindStringIndex(input)

	if len(indexes) == 0 {
		return false, NoTokenOffset
	}

	index := indexes[0]
	length := indexes[1] - index

	return true, TokenOffset{index, length}
}
