package ast

type RuleConcatenation struct {
	Rule
	Rules RuleSet
}

func (r RuleConcatenation) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	var nodes NodeSequence
	var lexemes LexemeSequence

	remaining := sequence

	for _, rule := range r.Rules {
		matched, newremaining, child, err := rule.Match(grammar, remaining)
		remaining = newremaining

		if matched == false {
			return false, remaining, nil, err
		}

		if child.IsEmpty() {
			continue
		}

		if child.GetRuleIdentity() == RuleIdentityNone {
			if child.CountLexemeSequence() == 1 {
				lexemes = append(lexemes, child.GetLexemeSequence()[0])
			} else {
				// This is probably an optional
				nodes = append(nodes, child)
			}
		} else {
			nodes = append(nodes, child)
		}
	}

	node := Node{
		Rule:    r.GetIdentity(),
		Lexemes: lexemes,
		Nodes:   nodes,
	}

	return true, remaining, node, nil
}
