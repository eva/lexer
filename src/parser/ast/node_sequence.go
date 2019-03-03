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
	if null, instanceof := node.(NodeNull); instanceof {
		if null.IsEmpty() {
			return sequence
		}

		return sequence.Merge(null.GetNodeSequence())
	}

	return append(sequence, node)
}

// Merge will append the given node sequence to the end of this node sequence.
// This function uses the Add method so the same rules apply.
func (sequence NodeSequence) Merge(nodes NodeSequence) NodeSequence {
	this := sequence

	for _, node := range nodes {
		// Skip all null nodes.
		// This is because adding a null node should not allow cascading of its sequence.
		// At least this is the case right now, I haven't got a case to allow this.
		if _, instanceof := node.(NodeNull); instanceof {
			continue
		}

		this = this.Add(node)
	}

	return this
}
