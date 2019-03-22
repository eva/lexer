package json

import (
	"fmt"
	"testing"

	"../../parser"
	"../../parser/ast"
)

func TestParseGrammar_BasicLiteral(test *testing.T) {
	node, err := parser.ParseAny(Grammar, `123`)

	if err != nil {
		test.Errorf(`Was not expecting to get error: %#v`, err)
		return
	}

	if node.GetNodeType() != ast.NodeTypeRule {
		test.Errorf(`Expected returned node to be of type ast.NodeTypeRule, instead got: %T`, node)
		return
	}

	literal, literalcasted := node.(ast.NodeRule)

	if literalcasted == false {
		test.Error(`Expected returned node to be an instance of ast.NodeRule`)
		return
	}

	if literal.GetRuleIdentity() != RuleGrammar {
		test.Errorf(`Expected returned rule to be %#v, instead got %#v`, RuleGrammar, literal.GetRuleIdentity())
		return
	}

	if literal.Count() != 1 {
		test.Error(`Expected literal count to be one.`)
	}
}

func TestParseGrammar_ObjectWithoutPair(test *testing.T) {
	found, rule := Grammar.FindRule(RuleObject)

	if found == false {
		test.Error(`Expected to be able to find rule in grammar`)
		return
	}

	sequence := ast.LexemeSequence{
		ast.Lexeme{Token: TokenSyntaxCurlyBraceOpen},
		ast.Lexeme{Token: TokenSyntaxCurlyBraceClose},
	}

	matched, _, _, _ := rule.Match(Grammar, sequence)

	if matched == false {
		test.Error(`Expected to match`)
	}
}

func TestParseGrammar_ObjectWithSinglePair(test *testing.T) {
	found, rule := Grammar.FindRule(RuleObject)

	if found == false {
		test.Error(`Expected to be able to find rule in grammar`)
		return
	}

	offset := ast.TokenOffset{1, 1}
	sequence := ast.LexemeSequence{
		ast.Lexeme{Token: TokenSyntaxCurlyBraceOpen, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenSyntaxQuoteDouble, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenLiteralString, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenSyntaxQuoteDouble, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenSyntaxColon, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenLiteralBooleanTrue, Offset: offset, Value: `1`},
		ast.Lexeme{Token: TokenSyntaxCurlyBraceClose, Offset: offset, Value: `1`},
	}

	matched, _, node, _ := rule.Match(Grammar, sequence)

	if matched == false {
		test.Error(`Expected to match`)
		return
	}

	castee, instanceof := node.(ast.NodeRule)

	if instanceof == false {
		test.Error(`Expected to be a rule node`)
	}

	if castee.Count() == 2 {
		test.Error(`Expected node sequence to be greater than 2`)
	}

	fmt.Printf(`%#v`, node)
}
