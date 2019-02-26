package ast

type RuleOptional struct {
	Rule
	Target RuleKind
}

func (rule RuleOptional) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if sequence.IsEmpty() {
		err := NewErrRuleSequenceEmpty(rule)
		return false, sequence, nil, err
	}

	matched, remaining, node, err := rule.Target.Match(grammar, sequence)

	if err != nil {
		switch err.(type) {
		case *ErrRuleReferenceNotFound:
			return false, remaining, nil, err
		}
	}

	return matched, remaining, node, nil
}
