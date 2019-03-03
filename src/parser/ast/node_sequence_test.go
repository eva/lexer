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

func TestNodeSequence_AddNodeNullWithSequence(test *testing.T) {
	node := NodeNull{
		Nodes: NodeSequence{
			NodeLexeme{},
			NodeLexeme{},
			NodeLexeme{},
		},
	}

	sequence := NodeSequence{}
	returned := sequence.Add(node)

	if returned.IsEmpty() == true {
		test.Error(`Expected returned sequence not be empty, null nodes should merge there sequences`)
	}

	if returned.Count() != 3 {
		test.Errorf(`Expected returned sequence to have 3 count, got %d`, returned.Count())
	}
}

func TestNodeSequence_MergeNodeSequence(test *testing.T) {
	nodes := NodeSequence{
		NodeLexeme{Token: 1},
		NodeLexeme{Token: 2},
		NodeLexeme{Token: 3},
	}

	sequence := NodeSequence{
		NodeLexeme{Token: 4},
		NodeLexeme{Token: 5},
	}

	returned := sequence.Merge(nodes)

	if returned.IsEmpty() == true {
		test.Error(`Expected returned sequence not to be empty`)
	}

	if returned.Count() != 5 {
		test.Errorf(`Expected returned sequence to have 5 count, got %d`, returned.Count())
	}

	if castee, instanceof := returned[0].(NodeLexeme); instanceof && castee.GetTokenIdentity() != 4 {
		test.Error(`Expected the first node to remain the first node`)
	}

	if castee, instanceof := returned[2].(NodeLexeme); instanceof && castee.GetTokenIdentity() != 1 {
		test.Error(`Expected the third node to be the first node in the merge sequence`)
	}
}
