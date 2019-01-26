package ast

type NamespaceIdentity string

var NamespaceIdentityNone NamespaceIdentity = ""
var NamespaceIdentityShift NamespaceIdentity = "<<"

type NamespaceSet []NamespaceKind
type NamespaceStack NamespaceSet

type NamespaceKind interface {
	GetIdentity() NamespaceIdentity
	GetToken(id TokenIdentity) (found bool, token TokenKind)
	GetTokens() TokenSet
}

type Namespace struct {
	Identity NamespaceIdentity
	Tokens   TokenSet
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

func (n Namespace) GetTokens() TokenSet {
	return n.Tokens
}
