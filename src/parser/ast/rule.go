package ast

// RuleIdentity represents a unique identity given to all rules.
type RuleIdentity uint

// RuleIdentityNone is the zero token identity.
// The lower bound for a valid rule identity is 1, meaning zero (0) is invalid.
// When a rule identity is not specifically given an identity we can match against this variable.
var RuleIdentityNone RuleIdentity

// RuleSet is simply a collection of rule kind.
// This is defined simply as a shortcut and syntactical sugar when defining grammars.
type RuleSet []RuleKind

// RuleKind is an interface that vaguely wraps the core functionality for lexical rules.
// The core functionality for a rule is the ability to know its identity and to match against a lexeme sequence.
type RuleKind interface {
	HasIdentity() bool
	GetIdentity() RuleIdentity
	ShouldIgnore() bool
	Match(grammar GrammarKind, sequence LexemeSequence) (matched bool, remaining LexemeSequence, node NodeKind, err error)
}

// Rule is a basic core implementation for `RuleKind` minus the match method.
type Rule struct {
	Identity RuleIdentity
	Ignore   bool
}

// HasIdentity will validate that the rule has been initialised with a valid identity.
// An identity is valid if its greater than the uninitialised value.
func (rule Rule) HasIdentity() bool {
	return rule.GetIdentity() != RuleIdentityNone
}

// GetIdentity returns the rule identity the rule was initialised with.
// In cases where one wasn't provided at initialisation then `RuleIdentityNone` is returned.
// Make sure to check against `HasIdentity()` first.
func (rule Rule) GetIdentity() RuleIdentity {
	return rule.Identity
}

// ShouldIgnore returns if the rule should be ignored from root level parsing.
func (rule Rule) ShouldIgnore() bool {
	return rule.Ignore
}
