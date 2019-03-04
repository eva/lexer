package ast

// GrammarKind is an interface that represents a grammar at its highest level.
// A grammar represents a series of tokens and rules that can be used to parse source code.
type GrammarKind interface {
	FindRootNamespace() (found bool, namespace NamespaceKind)
	FindNamespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	FindRule(id RuleIdentity) (found bool, rule RuleKind)
	GetRules() RuleCollection
	IsTokenIgnored(id TokenIdentity) bool
}

// Grammar is a base implementing of GrammarKind.
type Grammar struct {
	Namespaces   NamespaceCollection
	Rules        RuleCollection
	IgnoreTokens TokenIdentityCollection
}

// FindRootNamespace is a shortcut method to find a namespace with the root identity.
func (grammar Grammar) FindRootNamespace() (bool, NamespaceKind) {
	return grammar.FindNamespace(NamespaceIdentityRoot)
}

// FindNamespace will attempt to find and return a namespace by id.
// When the identity is not found then false is returned with a nil namespace.
// When the namespace is found then true is returned with the found namepsace.
func (grammar Grammar) FindNamespace(id NamespaceIdentity) (bool, NamespaceKind) {
	for _, namespace := range grammar.Namespaces {
		if namespace.GetIdentity() == id {
			return true, namespace
		}
	}

	return false, nil
}

// FindRule will attempt to find and return a rule by its identity.
// When the identity is not found then false is returned with a nil rule.
// When the rule is found then true is returned with the found rule.
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

// IsTokenIgnored will return true if the token is mentioned in the ignore token collection.
// This denotes a token that is not important information to store in the parsed tree.
// For example whitespace in most languages is nothing but for the developers experience.
func (grammar Grammar) IsTokenIgnored(id TokenIdentity) bool {
	for _, token := range grammar.IgnoreTokens {
		if token == id {
			return true
		}
	}

	return false
}
