package ast

import "testing"

func TestNodeSequence_EmptySequence(test *testing.T) {
	sequence := NodeSequence{}

	if sequence.IsEmpty() != true {
		test.Error(`Expected empty sequence to detect with IsEmpty()`)
	}

	if sequence.Count() != 0 {
		test.Errorf(`Expected empty sequence to have 0 count, got %d`, sequence.Count())
	}
}

func TestNodeSequence_AddNodeRule(test *testing.T) {
	node := NodeRule{}

	sequence := NodeSequence{}
	returned := sequence.Add(node)

	if returned.IsEmpty() != false {
		test.Error(`Expected returned sequence to not be empty`)
	}

	if returned.Count() != 1 {
		test.Error(`Expected returned sequence to contain a new node`)
	}
}

func TestNodeSequence_AddNodeLexeme(test *testing.T) {
	node := NodeLexeme{}

	sequence := NodeSequence{}
	returned := sequence.Add(node)

	if returned.IsEmpty() != false {
		test.Error(`Expected returned sequence to not be empty`)
	}

	if returned.Count() != 1 {
		test.Error(`Expected returned sequence to contain a new node`)
	}
}

func TestNodeSequence_AddNodeNull(test *testing.T) {
	node := NodeNull{}

	sequence := NodeSequence{}
	returned := sequence.Add(node)

	if returned.IsEmpty() != true {
		test.Error(`Expected returned sequence to be empty when adding ast.NodeNull (which should be ignored)`)
	}

	if returned.Count() != 0 {
		test.Errorf(`Expected returned sequence to have 0 count, got %d`, returned.Count())
	}
}
