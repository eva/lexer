package ast

import "testing"

func TestLexemeSequence_EmptySequence(test *testing.T) {
	sequence := LexemeSequence{}

	if sequence.Count() != 0 {
		test.Errorf(`The sequence have a count of zero, got %d`, sequence.Count())
		return
	}

	if !sequence.IsEmpty() {
		test.Error(`The sequence with count zero should be considered empty`)
		return
	}
}

func TestLexemeSequence_WithEntries(test *testing.T) {
	sequence := LexemeSequence{
		Lexeme{},
		Lexeme{},
	}

	if sequence.Count() != 2 {
		test.Errorf(`Expected sequence to have a count of 2, instead got %d`, sequence.Count())
		return
	}

	if sequence.IsEmpty() != false {
		test.Error(`Expected sequence to not be empty with entries`)
		return
	}
}
