package ast

type RuleToken struct {
	Rule
	Target TokenIdentity
}

func (r RuleToken) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if sequence.IsEmpty() {
		return r.Rule.Match(grammar, sequence)
	}

	lexeme := sequence[0]

	if lexeme.IsTokenIdentity(r.Target) == false {
		return r.Rule.Match(grammar, sequence)
	}

	node := Node{
		Rule:    r.GetIdentity(),
		Lexemes: LexemeSequence{lexeme},
	}

	remaining := sequence[1:]

	return true, remaining, node, nil
}
