package ast

type RuleOptional struct {
	Rule
	Target RuleKind
}

func (r RuleOptional) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if sequence.IsEmpty() {
		err := &ErrRuleSequenceEmpty{
			RuleIdentity: r.GetIdentity(),
		}

		return false, sequence, nil, err
	}

	matched, remaining, node, err := r.Target.Match(grammar, sequence)

	if err != nil {
		switch err.(type) {
		case *ErrRuleReferenceNotFound:
			return false, remaining, nil, err
		}
	}

	return matched, remaining, node, nil
}
