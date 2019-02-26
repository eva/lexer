package ast

import "testing"

func TestNodeLexeme_GetNodeType(test *testing.T) {
	node := NodeLexeme{}

	if node.GetNodeType() != NodeTypeLexeme {
		test.Errorf(`Expected node type to be ast.NodeTypeLexeme, instead got %#v`, node.GetNodeType())
	}
}

func TestNodeLexeme_IsValid(test *testing.T) {
	a := NodeLexeme{}

	if a.IsValid() == true {
		test.Error(`Node initialised with all default values should be considered invalid`)
	}

	b := NodeLexeme{
		Token: 1,
	}

	if b.IsValid() == true {
		test.Error(`Node initialised with only a token identity should be considered invalid`)
	}

	c := NodeLexeme{
		Token:  1,
		Offset: TokenOffset{0, 1},
	}

	if c.IsValid() == true {
		test.Error(`Node initialised with missing value should be considered invalid`)
	}

	d := NodeLexeme{
		Token:  1,
		Offset: TokenOffset{0, 1},
		Value:  "a",
	}

	if d.IsValid() == false {
		test.Error(`Node initialised with valid data should be considered valid`)
	}
}

func TestNodeLexeme_BasicGetterFunctionality(test *testing.T) {
	node := NodeLexeme{
		Token:  1,
		Offset: TokenOffset{2, 3},
		Value:  "a",
	}

	if node.GetTokenIdentity() != 1 {
		test.Errorf(`Expected node token identity to be 1, instead got: %d`, node.GetTokenIdentity())
	}

	offset := node.GetTokenOffset()

	if offset[0] != 2 && offset[1] != 3 {
		test.Errorf(`Expected node offset to be [2,3], instead got: %#v`, offset)
	}

	if node.GetValue() != "a" {
		test.Errorf(`Expected node value to be "a", instead got: %#v`, node.GetValue())
	}
}
