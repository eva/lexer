package ast

type RuleToken struct {
	Rule
	Target TokenIdentity
}

func (r RuleToken) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if sequence.IsEmpty() {
		err := NewErrRuleSequenceEmpty(r)
		return false, sequence, nil, err
	}

	lexeme := sequence[0]

	if lexeme.IsTokenIdentity(r.Target) == false {
		err := NewErrRuleTokenMatchFailure(r, r.Target, lexeme)
		return false, sequence, nil, err
	}

	node := Node{
		Rule:    r.GetIdentity(),
		Lexemes: LexemeSequence{lexeme},
	}

	remaining := sequence[1:]

	return true, remaining, node, nil
}
