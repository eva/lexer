package ast

type NodeSequence []NodeKind

func (sequence NodeSequence) Count() int {
	return len(sequence)
}

func (sequence NodeSequence) IsEmpty() bool {
	return sequence.Count() == 0
}
