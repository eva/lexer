package ast

import "errors"

// ErrGrammarMissingRootNamespace is an error that is returned from the RootNamespace() function
// when there are no namespaces that qualify to be a root namespace.
var ErrGrammarMissingRootNamespace = errors.New("Could not resolve the root namespace for grammar")

type GrammarKind interface {
	RootNamespace() (namespace NamespaceKind, err error)
	Namespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	Token(id TokenIdentity) (found bool, token TokenKind)
	Rule(id RuleIdentity) (found bool, rule RuleKind)
	GetRules() RuleSet
}

type Grammar struct {
	Namespaces NamespaceSet
	Tokens     TokenSet
	Rules      RuleSet
}

func (g Grammar) RootNamespace() (NamespaceKind, error) {
	for _, namespace := range g.Namespaces {
		return namespace, nil
	}

	return nil, ErrGrammarMissingRootNamespace
}

func (g Grammar) Namespace(id NamespaceIdentity) (bool, NamespaceKind) {
	for _, namespace := range g.Namespaces {
		if namespace.GetIdentity() == id {
			return true, namespace
		}
	}

	return false, nil
}

func (g Grammar) Token(id TokenIdentity) (bool, TokenKind) {
	return false, nil
}

func (g Grammar) Rule(id RuleIdentity) (bool, RuleKind) {
	return false, nil
}

func (g Grammar) GetRules() RuleSet {
	return g.Rules
}
