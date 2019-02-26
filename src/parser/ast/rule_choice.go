package ast

type RuleChoice struct {
	Rule
	Rules []RuleKind
}

func (r RuleChoice) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	for _, rule := range r.Rules {
		matched, remaining, child, err := rule.Match(grammar, sequence)

		if err != nil {
			switch err.(type) {
			case *ErrRuleReferenceNotFound:
				return false, remaining, nil, err
			}
		}

		if matched == false {
			continue
		}

		node := NodeRule{
			Rule:  r.GetIdentity(),
			Nodes: NodeSequence{child},
		}

		return true, remaining, node, nil
	}

	err := NewErrRuleChoiceNoneMatched(r)
	return false, sequence, nil, err
}
