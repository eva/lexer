package ast

import (
	"regexp"
	"testing"
)

func TestTokenExpression_IsTokenKind(test *testing.T) {
	var token interface{} = TokenExpression{}
	_, instanceof := token.(TokenKind)

	if instanceof == false {
		test.Error(`Expected TokenExpression to be of TokenKind`)
		return
	}
}

func TestTokenExpression_Match(test *testing.T) {
	dataset := []struct {
		input   string
		value   string
		matched bool
		offset  TokenOffset
	}{
		{"foo", `o{1}`, true, TokenOffset{1, 1}},
		{"foo", `o{2}`, true, TokenOffset{1, 2}},
		{"foo", `o{3}`, false, NoTokenOffset},
		{"foo", `foop?`, true, TokenOffset{0, 3}},
	}

	for i, data := range dataset {
		i = i + 1

		token := TokenExpression{
			Expression: regexp.MustCompile(data.value),
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

func BenchmarkTokenExpressionMatch_MatchWord(b *testing.B) {
	token := TokenExpression{
		Expression: regexp.MustCompile(`foo`),
	}

	for i := 0; i < b.N; i++ {
		token.Match(`foo`)
	}
}

func BenchmarkTokenExpressionMatch_MatchWordSplit(b *testing.B) {
	token := TokenExpression{
		Expression: regexp.MustCompile(`foo|bar`),
	}

	for i := 0; i < b.N; i++ {
		token.Match(`foo`)
	}
}

func BenchmarkTokenExpressionMatch_MatchExpression(b *testing.B) {
	token := TokenExpression{
		Expression: regexp.MustCompile(`f?o+(ba|bu)r{1,2}`),
	}

	for i := 0; i < b.N; i++ {
		token.Match(`fooburr`)
	}
}
