package ast

import "errors"

// A RuleIdentity represents a rule identity in lexical parsing.
// These rule identities must be unique to each grammar.
type RuleIdentity int

var RuleIdentityNone RuleIdentity = 0

type RuleSet []RuleKind

// RuleKind is an interface that vaguely wraps the core functionality for lexical rules.
// A rule should contain an identity and be able to see if it matches a lexeme sequence.
type RuleKind interface {
	GetIdentity() RuleIdentity
	ShouldIgnore() bool
	Match(grammar GrammarKind, sequence LexemeSequence) (matched bool, remaining LexemeSequence, node NodeKind, err error)
}

var ErrRuleNotMatched = errors.New("The rule did not match")

type Rule struct {
	Identity RuleIdentity
	Ignore   bool
}

func (r Rule) GetIdentity() RuleIdentity {
	return r.Identity
}

func (r Rule) ShouldIgnore() bool {
	return r.Ignore
}

func (r Rule) Match(grammar GrammarKind, sequence LexemeSequence) (bool, LexemeSequence, NodeKind, error) {
	return false, sequence, nil, ErrRuleNotMatched
}
