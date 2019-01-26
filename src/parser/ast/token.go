package ast

// A TokenOffset represents a set of numbers indicating the start and end of a token value.
// Typically this will be used as a response from token matching as the value may not be important.
type TokenOffset [2]int

// NoTokenOffset represents a token that was not found and will be acompanied with a false match.
// All valid tokens should be of length 1, therefore {0,0} can be considered a "match miss".
var NoTokenOffset = TokenOffset{0, 0}

// A TokenIdentity represents a lexical token.
// The token identities must be unique per grammar.
type TokenIdentity int

// InvalidTokenIdentity represents an invalid token.
// Zero is reserved for cases where an invalid token identity is needed.
const InvalidTokenIdentity TokenIdentity = 0

type TokenIdentitySet []TokenIdentity

// A TokenSet represents a set of TokenKind.
type TokenSet []TokenKind

// A TokenKind is a kind of lexical token.
type TokenKind interface {
	GetIdentity() TokenIdentity
	HasTransition() (should bool, namespace NamespaceIdentity)
	Match(input string) (matched bool, offset TokenOffset)
}

// A Token is basic implemention of TokenKind.
// This should be used in composition to satisfy the TokenKind interface but the TokenKind.Match()
// should be implemented against with that tokens functionality.
type Token struct {
	Identity     TokenIdentity
	TransitionTo NamespaceIdentity
}

func (t Token) GetIdentity() TokenIdentity {
	return t.Identity
}

// HasTransition implements TokenKind.HasTransition()
// This method will return true for should when the token should transition to another namespace.
// In this case the namespace will be NamespaceIdentity that can be looked up against the grammar.
// The other case is when a transition should not be made, indicating the next token should be
// taken from the current active namespace.
func (t Token) HasTransition() (should bool, namespace NamespaceIdentity) {
	if t.TransitionTo == NamespaceIdentityNone {
		return false, NamespaceIdentityNone
	}

	return true, t.TransitionTo
}

// Match implements TokenKind.Match()
// This method is hardcoded to always fail and return a NoTokenOffset as this method should be
// override by the struct that uses this in composition. Rembmer the Token struct is not a
// legal type of TokenKind.
func (t Token) Match(input string) (matched bool, offset TokenOffset) {
	return false, NoTokenOffset
}
