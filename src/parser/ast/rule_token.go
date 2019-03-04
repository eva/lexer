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
	id := rule.Target

	if sequence.IsEmpty() {
		err := NewErrRuleSequenceEmpty(rule)
		return false, sequence, nil, err
	}

	lexeme := sequence[0]

	if lexeme.IsTokenIdentity(id) == false {
		err := NewErrRuleTokenMatchFailure(rule, id, lexeme)
		return false, sequence, nil, err
	}

	remaining := sequence[1:]

	// The token can be ignored from the grammar.
	// In this case the match was a sucess but a null node is returned.
	if grammar.IsTokenIgnored(id) {
		node := NodeNull{}
		return true, remaining, node, nil
	}

	node := NewLexemeNode(lexeme)
	return true, remaining, node, nil
}
