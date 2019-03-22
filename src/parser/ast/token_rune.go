package ast

// TokenRune represents a single character that can be identified.
// This is benchmarked to be around twice as fast as TokenLiteral.
type TokenRune struct {
	Token
	Rune rune
}

// Match will attempt to take the first rune from the given input and match against the token's rune.
// Note this token match will only check the first rune and no more.
func (token TokenRune) Match(input string) (bool, TokenOffset) {
	r := rune(input[0])

	if r != token.Rune {
		return false, NoTokenOffset
	}

	length := len(string(token.Rune))
	offset := TokenOffset{0, length}

	return true, offset
}
