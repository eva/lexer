package ast

type Grammar interface {
}

// type Namespace interface {
// 	Identity() string
// 	Tokens() TokenSet
// }

type NamespaceSet []Namespace

// An TokenOffset represents a set of numbers indicating the start and end of a token value.
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

// A TokenSet represents a set tokens.
type TokenSet []Token

// A Token is a generic interface a kind of token.
type Token interface {
	Identity() TokenIdentity
	Namespace() Namespace
	Transition() (change bool, namespace string)
	Match(input string) (matched bool, offset TokenOffset)
}
