package ast

type RuleOptional struct {
	Rule
	Target RuleKind
}

func (r RuleOptional) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	var nodes NodeSequence

	matched, remaining, child, _ := r.Target.Match(grammar, sequence)

	if matched == true {
		nodes = append(nodes, child)
	}

	node := Node{
		Rule:  r,
		Nodes: nodes,
	}

	return true, remaining, node, nil
}
