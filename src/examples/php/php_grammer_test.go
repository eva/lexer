package php

import (
	"testing"

	"../../parser"
	"../../parser/ast"
)

func TestParseGrammar_BasicVariable(test *testing.T) {
	node, err := parser.ParseAny(Grammar, `$foo`)

	if err != nil {
		test.Errorf(`Was not expecting to get error: %v`, err)
		return
	}

	if node.GetNodeType() != ast.NodeTypeRule {
		test.Errorf(`Expected returned node to be of type ast.NodeTypeRule, instead got: %T`, node)
		return
	}

	cast, casted := node.(ast.NodeRule)

	if casted == false {
		test.Error(`Expected returned node to be an instance of ast.NodeRule`)
		return
	}

	if cast.GetRuleIdentity() != RuleVariable {
		test.Errorf(`Expected returned rule to be %#v, instead got %#v`, RuleVariable, cast.GetRuleIdentity())
		return
	}

	// log.Printf(`%#v`, node)
}

func TestParseGrammar_BasicExpression(test *testing.T) {
	node, err := parser.ParseAny(Grammar, `$foo - $bar`)

	if err != nil {
		test.Errorf(`Was not expecting to get error: %v`, err)
		return
	}

	if node.GetNodeType() != ast.NodeTypeRule {
		test.Errorf(`Expected returned node to be of type ast.NodeTypeRule, instead got: %T`, node)
		return
	}

	cast, casted := node.(ast.NodeRule)

	if casted == false {
		test.Error(`Expected returned node to be an instance of ast.NodeRule`)
		return
	}

	if cast.GetRuleIdentity() != RuleExpression {
		test.Errorf(`Expected returned rule to be %#v, instead got %#v`, RuleExpression, cast.GetRuleIdentity())
		return
	}

	// log.Printf(`%#v`, node)
}

func TestParseGrammar_BasicExpressionSelfReference(test *testing.T) {
	node, err := parser.ParseAny(Grammar, `1+1*1`)

	if err != nil {
		test.Errorf(`Was not expecting to get error: %v`, err)
		return
	}

	if node.GetNodeType() != ast.NodeTypeRule {
		test.Errorf(`Expected returned node to be of type ast.NodeTypeRule, instead got: %T`, node)
		return
	}

	cast, casted := node.(ast.NodeRule)

	if casted == false {
		test.Error(`Expected returned node to be an instance of ast.NodeRule`)
		return
	}

	if cast.GetRuleIdentity() != RuleExpression {
		test.Errorf(`Expected returned rule to be %#v, instead got %#v`, RuleExpression, cast.GetRuleIdentity())
		return
	}

	// log.Printf(`%#v`, node)
}

func TestTokeniseGrammar(test *testing.T) {
	dataset := []struct {
		input  string
		tokens []ast.TokenIdentity
	}{
		{`+`, []ast.TokenIdentity{TokenAddition}},
		{`+ +`, []ast.TokenIdentity{TokenAddition, TokenWhitespace, TokenAddition}},
		{`( + ]`, []ast.TokenIdentity{TokenSyntaxParenthesisOpen, TokenWhitespace, TokenAddition, TokenWhitespace, TokenSyntaxSquareBracketClose}},
		{`$foo`, []ast.TokenIdentity{TokenDollar, TokenIdentifier}},
	}

	for i, data := range dataset {
		i = i + 1

		sequence, index, err := parser.Tokenise(Grammar, data.input)

		if err != nil {
			test.Errorf(`[%d] Was not expecting to have an error: %v`, i, err)
			return
		}

		if index != len(data.input) {
			test.Errorf(`[%d] Was expecting to have consumed the entire input, index is %d`, i, index)
			return
		}

		if len(data.tokens) != len(sequence) {
			test.Errorf(`[%d] Expecting that length of sequence is %d but got %d`, i, len(data.tokens), len(sequence))
			return
		}

		for position, lexeme := range sequence {
			token := data.tokens[position]

			if lexeme.GetTokenIdentity() != token {
				test.Errorf(`[%d] @%d Token %v was expected, instead got %v`, i, position, token, lexeme.GetTokenIdentity())
				return
			}
		}
	}
}
