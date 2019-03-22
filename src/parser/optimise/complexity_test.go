package optimise

import (
	"testing"

	"../ast"
)

func TestCalculateComplexity_TokenLiteral(test *testing.T) {
	grammar := ast.Grammar{}
	token := ast.TokenLiteral{}

	result, err := calculateTokenComplexity(grammar, token)

	if err != nil {
		test.Errorf(`Was not expecting error: %+v`, err)
		return
	}

	if result != 1 {
		test.Error(`Expected ast.TokenListeral to have a single complexity`)
	}
}
