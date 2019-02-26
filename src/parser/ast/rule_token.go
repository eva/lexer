package ast

// RuleToken represents a rule that should match the target token identity.
type RuleToken struct {
	Rule
	Target TokenIdentity
}

// Match will attempt to match the target token against the given sequence.
// Assuming the sequence contains the target token in its first position the match is considered successful.
// The kind of node returned will always be an instance of ast.NodeLexeme.
func (rule RuleToken) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	if sequence.IsEmpty() {
		err := NewErrRuleSequenceEmpty(rule)
		return false, sequence, nil, err
	}

	lexeme := sequence[0]

	if lexeme.IsTokenIdentity(rule.Target) == false {
		err := NewErrRuleTokenMatchFailure(rule, rule.Target, lexeme)
		return false, sequence, nil, err
	}

	node := NodeLexeme{
		Token:  lexeme.GetTokenIdentity(),
		Offset: lexeme.GetTokenOffset(),
		Value:  lexeme.GetValue(),
	}

	remaining := sequence[1:]

	return true, remaining, node, nil
}
