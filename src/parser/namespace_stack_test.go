package parser

import (
	"testing"

	"./ast"
)

func TestNamespaceStack(test *testing.T) {
	stack := NamespaceStack{}

	var a ast.NamespaceKind
	var b ast.NamespaceKind

	if stack.Count() != 0 {
		test.Errorf(`Expected empty stack to have 0 count, instead got %d`, stack.Count())
		return
	}

	if stack.Current() != nil {
		test.Error(`Expected empty stack to not return current namespace`)
		return
	}

	stack.Register(a)

	if stack.Count() != 1 {
		test.Errorf(`Expected stack count to be 1, instead got %d`, stack.Count())
		return
	}

	if stack.Current() != a {
		test.Error(`Expected stack current to be the current namespace "a"`)
		return
	}

	stack.Register(b)

	if stack.Count() != 2 {
		test.Errorf(`Expected stack count to be 2, instead got %d`, stack.Count())
		return
	}

	if stack.Current() != b {
		test.Error(`Expected stack current to be the current namespace "b"`)
		return
	}

	stack.Shift()

	if stack.Count() != 1 {
		test.Errorf(`Expected stack count to be 1 after shift, instead got %d`, stack.Count())
		return
	}

	if stack.Current() != a {
		test.Error(`Expected stack current to be namespace "a" after shift`)
		return
	}
}
