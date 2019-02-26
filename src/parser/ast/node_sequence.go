package ast

// NodeSequence represents a sequence of nodes that have matched rules whilst parsing.
// The sequence will be a combination of node kinds and therefore will need to be type checked.
// Valid node kinds are rules and lexemes, null nodes are ignored.
// Note that node sequences are immutable and thus actions will always return new instances with the modifications.
type NodeSequence []NodeKind

// Count will return the amount of nodes within the sequence.
// This is simply a wrapper around the len() function.
func (sequence NodeSequence) Count() int {
	return len(sequence)
}

// IsEmpty will check if the sequence has an internal count of zero.
// This is simply a wrapper around the len() function and zero check.
func (sequence NodeSequence) IsEmpty() bool {
	return sequence.Count() == 0
}

// Add will attempt to add the given node kind to the sequence.
// Node kinds such as rule and lexeme are added.
// Others are ignored and the same sequence is returned instead.
func (sequence NodeSequence) Add(node NodeKind) NodeSequence {
	switch node.(type) {
	case NodeNull:
		return sequence
	}

	return append(sequence, node)
}
