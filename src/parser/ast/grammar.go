package ast

// GrammarKind is an interface that represents a grammar at its highest level.
// A grammar represents a series of tokens and rules that can be used to parse source code.
type GrammarKind interface {
	FindRootNamespace() (found bool, namespace NamespaceKind)
	FindNamespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	FindToken(id TokenIdentity) (found bool, token TokenKind)
	FindRule(id RuleIdentity) (found bool, rule RuleKind)
	GetRules() RuleSet
}

// Grammar is a fully implemented `GrammarKind`.
type Grammar struct {
	Namespaces NamespaceSet
	Tokens     TokenSet
	Rules      RuleSet
}

// FindRootNamespace is a shortcut method to find a namespace with the root identity.
func (g Grammar) FindRootNamespace() (bool, NamespaceKind) {
	return g.FindNamespace(NamespaceIdentityRoot)
}

// FindNamespace will attempt to find and return a namespace by id.
// When the identity is not found then false is returned with a nil namespace.
func (g Grammar) FindNamespace(id NamespaceIdentity) (bool, NamespaceKind) {
	for _, namespace := range g.Namespaces {
		if namespace.GetIdentity() == id {
			return true, namespace
		}
	}

	return false, nil
}

// FindToken will attempt to find and return a token by id.
// When the identity is not found then false is returned with a nil token.
func (g Grammar) FindToken(id TokenIdentity) (bool, TokenKind) {
	for _, token := range g.Tokens {
		if token.GetIdentity() == id {
			return true, token
		}
	}

	return false, nil
}

// FindRule will attempt to find and return a rule by id.
// When the identity is not found then false is returned with a nil rule.
func (g Grammar) FindRule(id RuleIdentity) (bool, RuleKind) {
	for _, rule := range g.Rules {
		if rule.GetIdentity() == id {
			return true, rule
		}
	}

	return false, nil
}

// GetRules will return all defined rules against the grammar.
func (g Grammar) GetRules() RuleSet {
	return g.Rules
}
