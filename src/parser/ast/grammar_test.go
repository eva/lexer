package ast

import "testing"

func TestGrammar_FindNamespace_WithNoNamespaces_MissingNamespace(test *testing.T) {
	grammar := Grammar{}

	found, namespace := grammar.FindNamespace("missing")

	if found != false {
		test.Error(`Expected to have not found a namespace`)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected to have a nil namespace but got: %v`, namespace)
		return
	}
}

func TestGrammar_FindNamespace_WithNamespaces_MissingNamespace(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{Identity: "foo"},
			Namespace{Identity: "bar"},
		},
	}

	found, namespace := grammar.FindNamespace("missing")

	if found != false {
		test.Error(`Expected to have not found a namespace`)
		return
	}

	if namespace != nil {
		test.Errorf(`Expected to have a nil namespace but got: %v`, namespace)
		return
	}
}

func TestGrammar_FindNamespace_WithNamespaces_FindNamespace(test *testing.T) {
	grammar := Grammar{
		Namespaces: NamespaceSet{
			Namespace{Identity: "foo"},
			Namespace{Identity: "bar"},
		},
	}

	found, namespace := grammar.FindNamespace("bar")

	if found != true {
		test.Error(`Expected to find namespace`)
		return
	}

	if namespace == nil {
		test.Error(`Expected to have namespace return when found is true`)
		return
	}

	if namespace.GetIdentity() != "bar" {
		test.Errorf(`Expected to have found namespace by identity, instead got: %v`, namespace.GetIdentity())
		return
	}
}
