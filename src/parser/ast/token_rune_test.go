package ast

import "testing"

func TestTokenRune_IsTokenKind(test *testing.T) {
	var token interface{} = TokenRune{}
	_, instanceof := token.(TokenKind)

	if instanceof == false {
		test.Error(`Expected TokenRune to be of TokenKind`)
		return
	}
}

func TestTokenRune_Match(test *testing.T) {
	dataset := []struct {
		id      int
		input   string
		value   rune
		matched bool
		offset  TokenOffset
	}{
		{1, "foo", 'b', false, NoTokenOffset},
		{2, "bar", 'f', false, NoTokenOffset},
		{3, "f", 'f', true, TokenOffset{0, 1}},
		{4, "foo", 'f', true, TokenOffset{0, 1}},
		{5, "foobar", 'b', false, NoTokenOffset},
		{6, "foobar", 'o', false, NoTokenOffset},
	}

	for _, data := range dataset {

		token := TokenRune{
			Rune: data.value,
		}

		matched, offset := token.Match(data.input)

		if matched != data.matched {
			test.Errorf(`[%d] Matched %v is expected to be %v`, data.id, matched, data.matched)
			return
		}

		if data.matched == false && offset != NoTokenOffset {
			test.Errorf(`[%d] The matched was expected to fail but an offset was returned still %v`, data.id, offset)
			return
		}

		if offset[0] != data.offset[0] || offset[1] != data.offset[1] {
			test.Errorf(`[%d] Offset %v was expected to match %v`, data.id, offset, data.offset)
			return
		}
	}
}

func BenchmarkTokenRuneMatch_Match(b *testing.B) {
	token := TokenRune{
		Rune: 'f',
	}

	for i := 0; i < b.N; i++ {
		token.Match(`foo`)
	}
}
