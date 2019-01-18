package ast

type GrammarKind interface {
	GetRootNamespace() NamespaceKind
	GetNamespace(id NamespaceIdentity) (found bool, namespace NamespaceKind)
	GetToken(id TokenIdentity) (found bool, token TokenKind)
	Validate() error
}

type Grammar struct {
	Namespaces NamespaceSet
	Tokens     TokenSet
}

func (g Grammar) GetRootNamespace() NamespaceKind {
	return nil
}

func (g Grammar) GetNamespace(id NamespaceIdentity) (found bool, namespace NamespaceKind) {
	return false, nil
}

func (g Grammar) GetToken(id TokenIdentity) (found bool, token TokenKind) {
	return false, nil
}

func (g Grammar) Validate() error {
	return nil
}
