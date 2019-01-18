package ast

import "testing"

func TestTokeLiteral_IsTokenKind(test *testing.T) {
	var token interface{} = TokenLiteral{}
	_, instanceof := token.(TokenKind)

	if instanceof == false {
		test.Error(`Expected TokenLiteral to be of TokenKind`)
		return
	}
}

func TestTokenLiteral_Match(test *testing.T) {
	dataset := []struct {
		input      string
		expression string
		matched    bool
		offset     TokenOffset
	}{
		{"foo", "bar", false, NoTokenOffset},
		{"bar", "bard", false, NoTokenOffset},
		{"fooo", "foooo", false, NoTokenOffset},
		{"foo", "foo", true, TokenOffset{0, 3}},
		{"foobar", "bar", true, TokenOffset{3, 3}},
		{"foobar", "foobar", true, TokenOffset{0, 6}},
		{"foofoob", "foo", true, TokenOffset{0, 3}},
		{"foofoob", "foob", true, TokenOffset{3, 4}},
	}

	for i, data := range dataset {
		i = i + 1

		token := TokenLiteral{
			Literal: data.expression,
		}

		matched, offset := token.Match(data.input)

		if matched != data.matched {
			test.Errorf(`[%d] Matched %v is expected to be %v`, i, matched, data.matched)
			return
		}

		if data.matched == false && offset != NoTokenOffset {
			test.Errorf(`[%d] The matched was expected to fail but an offset was returned still %v`, i, offset)
			return
		}

		if offset[0] != data.offset[0] || offset[1] != data.offset[1] {
			test.Errorf(`[%d] Offset %v was expected to match %v`, i, offset, data.offset)
			return
		}
	}
}
