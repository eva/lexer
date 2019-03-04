package ast

// GrammarKind is an interface that represents a grammar at its highest level.
// A grammar represents a series of tokens and rules that can be used to parse source code.
type GrammarKind interface {
	FindRootNamespace() (found bool, namespace NamespaceKind)
	FindNamespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	FindToken(id TokenIdentity) (found bool, token TokenKind)
	FindRule(id RuleIdentity) (found bool, rule RuleKind)
	GetRules() RuleCollection
}

// Grammar is a fully implemented `GrammarKind`.
type Grammar struct {
	Namespaces NamespaceCollection
	Tokens     TokenCollection // @todo Remove, can be implied from Namespaces
	Rules      RuleCollection
}

// FindRootNamespace is a shortcut method to find a namespace with the root identity.
func (grammar Grammar) FindRootNamespace() (bool, NamespaceKind) {
	return grammar.FindNamespace(NamespaceIdentityRoot)
}

// FindNamespace will attempt to find and return a namespace by id.
// When the identity is not found then false is returned with a nil namespace.
func (grammar Grammar) FindNamespace(id NamespaceIdentity) (bool, NamespaceKind) {
	for _, namespace := range grammar.Namespaces {
		if namespace.GetIdentity() == id {
			return true, namespace
		}
	}

	return false, nil
}

// FindToken will attempt to find and return a token by id.
// When the identity is not found then false is returned with a nil token.
func (grammar Grammar) FindToken(id TokenIdentity) (bool, TokenKind) {
	for _, token := range grammar.Tokens {
		if token.GetIdentity() == id {
			return true, token
		}
	}

	return false, nil
}

// FindRule will attempt to find and return a rule by id.
// When the identity is not found then false is returned with a nil rule.
func (grammar Grammar) FindRule(id RuleIdentity) (bool, RuleKind) {
	for _, rule := range grammar.Rules {
		if rule.GetIdentity() == id {
			return true, rule
		}
	}

	return false, nil
}

// GetRules will return all defined rules against the grammar.
func (grammar Grammar) GetRules() RuleCollection {
	return grammar.Rules
}
