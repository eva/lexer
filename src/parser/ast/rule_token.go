package ast

// RuleToken is an RuleKind that will match a single token.
// This can be considered the most basic of rules and will most likely be combined with other rules.
type RuleToken struct {
	Rule
	Target TokenIdentity
}

// MatchLexemeSequence will take the first lexeme in the sequence and attempt to match it against the internal token.
// When the first token in the sequence is not matched the match fails entirely.
// The node returned will always have a length of 1.
func (r RuleToken) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	lexeme := sequence[0]

	if lexeme.IsTokenIdentity(r.Target) == false {
		return r.Rule.Match(grammar, sequence)
	}

	node := Node{
		Rule:    r,
		Lexemes: LexemeSequence{lexeme},
	}

	remaining := sequence[1:]

	return true, remaining, node, nil
}
