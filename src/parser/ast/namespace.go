package ast

type NamespaceIdentity string

const NamespaceIdentityNone NamespaceIdentity = ""
const NamespaceIdentityRoot NamespaceIdentity = "root"
const NamespaceIdentityShift NamespaceIdentity = "<<"

type NamespaceSet []NamespaceKind
type NamespaceStack NamespaceSet

type NamespaceKind interface {
	GetIdentity() NamespaceIdentity
	GetToken(id TokenIdentity) (found bool, token TokenKind)
	GetTokens() TokenCollection
}

type Namespace struct {
	Identity NamespaceIdentity
	Tokens   TokenCollection
}

func (n Namespace) GetIdentity() NamespaceIdentity {
	return n.Identity
}

func (n Namespace) GetToken(id TokenIdentity) (found bool, token TokenKind) {
	for _, token := range n.Tokens {
		if token.GetIdentity() == id {
			return true, token
		}
	}

	return false, nil
}

func (n Namespace) GetTokens() TokenCollection {
	return n.Tokens
}
