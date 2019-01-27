package ast

type RuleReference struct {
	Rule
	Target RuleIdentity
}

func (r RuleReference) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	found, rule := grammar.FindRule(r.Target)

	if found == false {
		err := &ErrRuleReferenceNotFound{
			RuleIdentity: r.Target,
		}

		return false, sequence, nil, err
	}

	return rule.Match(grammar, sequence)
}
