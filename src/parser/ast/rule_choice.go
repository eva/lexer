package ast

type RuleChoice struct {
	Rule
	Rules []RuleKind
}

func (r RuleChoice) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	for _, rule := range r.Rules {
		matched, remaining, child, _ := rule.Match(grammar, sequence)

		if matched == false {
			continue
		}

		node := Node{
			Rule:  r.GetIdentity(),
			Nodes: NodeSequence{child},
		}

		return true, remaining, node, nil
	}

	return r.Rule.Match(grammar, sequence)
}
