package ast

// A TokenOffset represents a set of numbers indicating the start and end of a token value.
// Typically this will be used as a response from token matching as the value may not be important.
type TokenOffset [2]int

// IsValid will return true assuming the data the offset was constructed with is considered valud.
// That is the first integer representing the offset must be at least 0.
// The second number representing the length of the value captured should be greater or equal to 1.
func (offset TokenOffset) IsValid() bool {
	if offset[0] < 0 {
		return false
	}

	if offset[1] <= 0 {
		return false
	}

	return true
}

// NoTokenOffset represents a token that was not found and will be acompanied with a false match.
// All valid tokens should be of length 1, therefore {0,0} can be considered a "match miss".
var NoTokenOffset = TokenOffset{0, 0}

// A TokenIdentity represents a lexical token.
// The token identities must be unique per grammar.
type TokenIdentity int

// TokenIdentityNone represents an invalid token.
// Zero is reserved for cases where an invalid token identity is needed.
const TokenIdentityNone TokenIdentity = 0

type TokenIdentityCollection []TokenIdentity

type TokenName string

const TokenNameNone = ""

// A TokenCollection represents a series of TokenKind in a collection.
type TokenCollection []TokenKind

// A TokenKind is a kind of lexical token.
type TokenKind interface {
	GetIdentity() TokenIdentity
	GetName() TokenName
	HasTransition() (should bool, namespace NamespaceIdentity)
	Match(input string) (matched bool, offset TokenOffset)
}

// A Token is basic implemention of TokenKind.
// This should be used in composition to satisfy the TokenKind interface but the TokenKind.Match()
// should be implemented against with that tokens functionality.
type Token struct {
	Identity     TokenIdentity
	Name         TokenName
	TransitionTo NamespaceIdentity
}

// GetIdentity will return the given TokenIdentity.
func (token Token) GetIdentity() TokenIdentity {
	return token.Identity
}

// GetName will return the given TokenName.
func (token Token) GetName() TokenName {
	return token.Name
}

// HasTransition implements TokenKind.HasTransition()
// This method will return true for should when the token should transition to another namespace.
// In this case the namespace will be NamespaceIdentity that can be looked up against the grammar.
// The other case is when a transition should not be made, indicating the next token should be
// taken from the current active namespace.
func (token Token) HasTransition() (should bool, namespace NamespaceIdentity) {
	if token.TransitionTo == NamespaceIdentityNone {
		return false, NamespaceIdentityNone
	}

	return true, token.TransitionTo
}

// Match implements TokenKind.Match()
// This method is hardcoded to always fail and return a NoTokenOffset as this method should be
// override by the struct that uses this in composition. Rembmer the Token struct is not a
// legal type of TokenKind.
func (Token) Match(input string) (matched bool, offset TokenOffset) {
	return false, NoTokenOffset
}
