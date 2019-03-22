package optimise

import (
	"fmt"

	"../ast"
)

func CheckComplexity(grammar ast.GrammarKind, rule ast.RuleKind) (int, error) {
	return calculateRuleComplexity(grammar, rule)
}

func calculateTokenComplexity(grammar ast.GrammarKind, token ast.TokenKind) (int, error) {
	switch token.(type) {
	case ast.TokenLiteral:
		return 1, nil
	case ast.TokenRegex:
		return 5, nil
	}

	return 0, fmt.Errorf("Cannot define complexity of %T", token)
}

func calculateRuleComplexity(grammar ast.GrammarKind, rule ast.RuleKind) (int, error) {
	switch instance := rule.(type) {
	case ast.RuleToken:
		found, token := instance.FindToken(grammar)

		if found == false {
			return 1, nil
		}

		return calculateTokenComplexity(grammar, token)
	case ast.RuleOptional:
		complexity, err := calculateRuleComplexity(grammar, instance)

		if err != nil {
			return 0, err
		}

		complexity = complexity + 1

		return complexity, nil
	case ast.RuleRepetition:
		return 1, nil
	}

	return 0, fmt.Errorf("Cannot define complexity of %T", rule)
}
