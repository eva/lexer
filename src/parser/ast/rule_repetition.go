package ast

import "errors"

type RuleRepetition struct {
	Rule
	Target  RuleKind
	Minimum int
	Maximum int
}

func (rule RuleRepetition) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if rule.Maximum == 0 {
		err := NewErrRuleRepetitionMaximumZero(rule)
		return false, sequence, nil, err
	}

	matches := 0
	nodes := NodeSequence{}

	newremaining := sequence

	for {
		matched, remaining, node, err := rule.Target.Match(grammar, newremaining)
		newremaining = remaining

		if err != nil {
			switch err.(type) {
			case *ErrRuleReferenceNotFound:
				return false, remaining, nil, err
			}
		}

		if matched == false {
			break
		}

		matches++
		nodes = nodes.Add(node)

		// If the maximum is set to -1 it is indicating a open ended many match.
		// This means the loop can go forever /shrug
		if rule.Maximum == -1 {
			continue
		}

		if matches > rule.Maximum {
			break
		}
	}

	if matches < rule.Minimum {
		err := errors.New(`Minimum reach`)
		return false, sequence, nil, err
	}

	newnode := NewRuleNode(rule.GetIdentity(), nodes)

	return true, newremaining, newnode, nil
}
