package ast

type RuleOptional struct {
	Rule
	Target RuleKind
}

func (rule RuleOptional) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	matched, remaining, node, err := rule.Target.Match(grammar, sequence)

	if err != nil {
		switch err.(type) {
		case *ErrRuleReferenceNotFound:
			return false, remaining, nil, err
		}
	}

	if matched == false {
		return true, remaining, NodeNull{}, nil
	}

	return matched, remaining, node, nil
}
