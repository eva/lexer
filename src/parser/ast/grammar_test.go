package ast

import "testing"

func TestGrammar_GetRules_ReturnEmptyRuleset(test *testing.T) {
	grammar := Grammar{}
	rules := grammar.GetRules()

	if rules.IsEmpty() != true {
		test.Error(`Expected grammar initialised with defaults returns empty ruleset`)
	}
}

func TestGrammar_FindRootNamespace_WithNoNamespace(test *testing.T) {
	grammar := Grammar{}
	found, namespace := grammar.FindRootNamespace()

	if found == true {
		test.Errorf(`Expected to not find a root namespace, grammar is: %#v`, grammar)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected with "found=false" the namespace should be nil, got: %#v`, namespace)
		return
	}
}

func TestGrammar_FindRootNamespace_WithNoRootNamespace(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{},
			Namespace{},
			Namespace{},
		},
	}

	found, namespace := grammar.FindRootNamespace()

	if found == true {
		test.Errorf(`Expected to not find a root namespace as no namespaces have root identity, got: %#v`, namespace)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected with "found=false" the namespace should be nil, got: %#v`, namespace)
		return
	}
}

func TestGrammar_FindRootNamespace_CanFindRootNamespace(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{},
			Namespace{},
			Namespace{Identity: NamespaceIdentityRoot},
		},
	}

	found, namespace := grammar.FindRootNamespace()

	if found == false {
		test.Errorf(`Expected to have found root namespace, grammar is: %#v`, grammar)
		return
	}

	if namespace == nil {
		test.Error(`Expected with "found=true" the namespace should be returned`)
		return
	}
}

func TestGrammar_FindNamespace_WithNoNamespace(test *testing.T) {
	grammar := Grammar{}
	found, namespace := grammar.FindNamespace("foo")

	if found != false {
		test.Error(`Expected to have not found a namespace`)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected to have a nil namespace but got: %#v`, namespace)
		return
	}
}

func TestGrammar_FindNamespace_WithNamespaceWrongIdentity(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{Identity: "foo"},
			Namespace{Identity: "bar"},
		},
	}

	found, namespace := grammar.FindNamespace("baz")

	if found != false {
		test.Error(`Expected to have not found a namespace as identity is not defined`)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected to have a nil namespace but got: %#v`, namespace)
		return
	}
}

func TestGrammar_FindNamespace_CanMatchNamespace(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{Identity: "foo"},
			Namespace{Identity: "bar"},
		},
	}

	found, namespace := grammar.FindNamespace("bar")

	if found != true {
		test.Error(`Expected to find namespace as identity is defined`)
		return
	}

	if namespace == nil {
		test.Error(`Expected that when "found=true" a namespace should be returned also`)
		return
	}

	if namespace.GetIdentity() != "bar" {
		test.Errorf(`Expected to have found namespace by identity, instead got: %#v`, namespace)
		return
	}
}

func TestGrammar_FindToken_WithNoToken(test *testing.T) {
	grammar := Grammar{}
	found, token := grammar.FindToken(1)

	if found != false {
		test.Error(`Expected to have not found a token`)
		return
	}

	if token != nil {
		test.Errorf(`Expected to have a nil token but got: %#v`, token)
		return
	}
}

func TestGrammar_FindToken_WithTokenWrongIdentity(test *testing.T) {
	grammar := Grammar{
		Tokens: TokenSet{
			Token{Identity: 1},
			Token{Identity: 2},
		},
	}

	found, token := grammar.FindToken(3)

	if found != false {
		test.Error(`Expected to have not found a token as identity is not defined`)
		return
	}

	if token != nil {
		test.Errorf(`Expected to have a nil token but got: %#v`, token)
		return
	}
}

func TestGrammar_FindToken_CanMatchToken(test *testing.T) {
	grammar := Grammar{
		Tokens: TokenSet{
			Token{Identity: 1},
			Token{Identity: 2},
		},
	}

	found, token := grammar.FindToken(2)

	if found != true {
		test.Error(`Expected to find token as identity is defined`)
		return
	}

	if token == nil {
		test.Error(`Expected that when "found=true" a token should be returned also`)
		return
	}

	if token.GetIdentity() != 2 {
		test.Errorf(`Expected to have found token by identity, instead got: %#v`, token)
		return
	}
}

func TestGrammar_FindRule_WithNoRule(test *testing.T) {
	grammar := Grammar{}
	found, rule := grammar.FindRule(1)

	if found != false {
		test.Error(`Expected to have not found a rule`)
		return
	}

	if rule != nil {
		test.Errorf(`Expected to have a nil rule but got: %#v`, rule)
		return
	}
}

func TestGrammar_FindRule_WithRuleWrongIdentity(test *testing.T) {
	grammar := Grammar{
		Rules: RuleSet{
			RuleToken{Rule: Rule{Identity: 1}},
			RuleToken{Rule: Rule{Identity: 2}},
		},
	}

	found, rule := grammar.FindRule(3)

	if found != false {
		test.Error(`Expected to have not found a rule as identity is not defined`)
		return
	}

	if rule != nil {
		test.Errorf(`Expected to have a nil rule but got: %#v`, rule)
		return
	}
}

func TestGrammar_FindRule_CanMatchRule(test *testing.T) {
	grammar := Grammar{
		Rules: RuleSet{
			RuleToken{Rule: Rule{Identity: 1}},
			RuleToken{Rule: Rule{Identity: 2}},
		},
	}

	found, rule := grammar.FindRule(2)

	if found != true {
		test.Error(`Expected to find rule as identity is defined`)
		return
	}

	if rule == nil {
		test.Error(`Expected that when "found=true" a rule should be returned also`)
		return
	}

	if rule.GetIdentity() != 2 {
		test.Errorf(`Expected to have found rule by identity, instead got: %#v`, rule)
		return
	}
}
