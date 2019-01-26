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
	for _, rule := range grammar.GetRules() {
		matched, _, node, _ := rule.Match(grammar, sequence)

		if matched == true {
			return node, nil
		}
	}

	return nil, errors.New(`Failed to parse sequence (mode: any)`)
}
