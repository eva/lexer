package ast

import "errors"

// ErrGrammarMissingRootNamespace is an error that is returned from the RootNamespace() function
// when there are no namespaces that qualify to be a root namespace.
var ErrGrammarMissingRootNamespace = errors.New("Could not resolve the root namespace for grammar")

type GrammarKind interface {
	RootNamespace() (namespace NamespaceKind, err error)
	Namespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	Token(id TokenIdentity) (found bool, token TokenKind)
	Validate() error
}

type Grammar struct {
	Namespaces NamespaceSet
	Tokens     TokenSet
}

func (g Grammar) RootNamespace() (namespace NamespaceKind, err error) {
	for _, namespace := range g.Namespaces {
		return namespace, nil
	}

	return nil, ErrGrammarMissingRootNamespace
}

func (g Grammar) Namespace(id NamespaceIdentity) (found bool, namespace NamespaceKind) {
	return false, nil
}

func (g Grammar) Token(id TokenIdentity) (found bool, token TokenKind) {
	return false, nil
}

func (g Grammar) Validate() error {
	return nil
}
