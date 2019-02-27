package ast

// LexemeSequence is collection of lexeme kind.
type LexemeSequence []LexemeKind

func (sequence LexemeSequence) Count() int {
	return len(sequence)
}

func (sequence LexemeSequence) IsEmpty() bool {
	return len(sequence) == 0
}

func (sequence LexemeSequence) IsValid() bool {
	for _, lexeme := range sequence {
		if lexeme.IsValid() == false {
			return false
		}
	}

	return true
}
