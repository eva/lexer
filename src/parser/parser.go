package parser

import (
	"errors"

	"./ast"
)

func ParseAny(grammar ast.GrammarKind, input string) (ast.NodeKind, error) {
	sequence, _, _ := Tokenise(grammar, input)

	return ParseAnySequence(grammar, sequence)
}

func ParseAnySequence(grammar ast.GrammarKind, sequence ast.LexemeSequence) (ast.NodeKind, error) {
	rules := grammar.GetRules()

	for _, rule := range rules {
		if rule.ShouldIgnore() {
			continue
		}

		matched, sequence, node, _ := rule.Match(grammar, sequence)

		if matched == true {
			if sequence.IsEmpty() == false {
				continue
			}

			return node, nil
		}
	}

	return nil, errors.New(`Failed to parse sequence (mode: any)`)
}
