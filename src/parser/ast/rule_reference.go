package ast

type RuleReference struct {
	Rule
	Target RuleIdentity
}

func (r RuleReference) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	found, rule := grammar.FindRule(r.Target)

	if found == false {
		return r.Rule.Match(grammar, sequence)
	}

	return rule.Match(grammar, sequence)
}
