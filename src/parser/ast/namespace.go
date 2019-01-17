package ast

// A Namespace is a type of token set that limits the active tokens when parsing.
type Namespace struct {
	identity string
	tokens   TokenSet
}

func NewNamespace(identity string) Namespace {
	namespace := Namespace{}
	namespace.identity = identity

	return namespace
}

func NewNamespaceRoot() Namespace {
	return NewNamespace("root")
}

func (namespace Namespace) Identity() string {
	return namespace.identity
}

func (namespace Namespace) Tokens() TokenSet {
	return namespace.tokens
}
