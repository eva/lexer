package ast

type NamespaceIdentity string
type NamespaceSet []NamespaceKind
type NamespaceStack NamespaceSet

type NamespaceKind interface {
	GetTokens() TokenSet
	RegisterToken(token TokenKind) error
	Validate() error
}

type Namespace struct {
	Identity NamespaceIdentity
	Tokens   TokenSet
}

func (n Namespace) GetTokens() TokenSet {
	return n.Tokens
}

func (n Namespace) RegisterToken(token TokenKind) error {
	return nil
}

func (n Namespace) Validate() error {
	return nil
}
