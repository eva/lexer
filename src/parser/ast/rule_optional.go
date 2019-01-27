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

	var nodes NodeSequence

	matched, remaining, child, err := r.Target.Match(grammar, sequence)

	if err != nil {
		switch err.(type) {
		case *ErrRuleReferenceNotFound:
			return false, remaining, nil, err
		}
	}

	if matched == true {
		nodes = append(nodes, child)
	}

	node := Node{
		Rule:  r.GetIdentity(),
		Nodes: nodes,
	}

	return true, remaining, node, nil
}
