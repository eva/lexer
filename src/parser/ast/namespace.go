package ast

type NamespaceIdentity string

const NamespaceIdentityNone NamespaceIdentity = ""
const NamespaceIdentityRoot NamespaceIdentity = "root"
const NamespaceIdentityShift NamespaceIdentity = "<<"

type NamespaceCollection []NamespaceKind
type NamespaceStack NamespaceCollection

type NamespaceKind interface {
	GetIdentity() NamespaceIdentity
	GetTokens() TokenCollection
	FindToken(id TokenIdentity) (found bool, token TokenKind)
}

type Namespace struct {
	Identity NamespaceIdentity
	Tokens   TokenCollection
}

func (n Namespace) GetIdentity() NamespaceIdentity {
	return n.Identity
}

func (n Namespace) GetTokens() TokenCollection {
	return n.Tokens
}

func (n Namespace) FindToken(id TokenIdentity) (found bool, token TokenKind) {
	for _, token := range n.Tokens {
		if token.GetIdentity() == id {
			return true, token
		}
	}

	return false, nil
}
