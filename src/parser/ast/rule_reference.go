package ast

// RuleReference represents a reference to another rule available within the grammar.
// The reference rule will then be matched against instead, this acting as a proxy for that rule.
type RuleReference struct {
	Rule
	Target RuleIdentity
}

// Match will attempt to find the target rule against the grammar and match against that instead.
// The returned node kind is directly passed from the proxied rule, therefore this rule is not represented as a node.
func (rule RuleReference) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	found, proxy := grammar.FindRule(rule.Target)

	if found == false {
		err := NewErrRuleReferenceNotFound(rule, rule.Target)
		return false, sequence, nil, err
	}

	return proxy.Match(grammar, sequence)
}
