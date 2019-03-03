package ast

import "fmt"

type ErrRuleSequenceEmpty struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleSequenceEmpty(rule RuleKind) error {
	return &ErrRuleSequenceEmpty{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: rule.GetIdentity(),
	}
}

func (e ErrRuleSequenceEmpty) Error() string {
	message := `The rule cannot match against an empty lexeme sequence.`

	return fmt.Sprintf(`[ErrRuleSequenceEmpty] %s`, message)
}

type ErrRuleTokenMatchFailure struct {
	RuleType             string
	RuleIdentity         RuleIdentity
	TargetTokenIdentity  TokenIdentity
	CurrentTokenIdentity TokenIdentity
	CurrentTokenOffset   TokenOffset
}

func NewErrRuleTokenMatchFailure(rule RuleKind, target TokenIdentity, lexeme LexemeKind) error {
	return &ErrRuleTokenMatchFailure{
		RuleType:             fmt.Sprintf(`%T`, rule),
		RuleIdentity:         rule.GetIdentity(),
		TargetTokenIdentity:  target,
		CurrentTokenIdentity: lexeme.GetTokenIdentity(),
		CurrentTokenOffset:   lexeme.GetTokenOffset(),
	}
}

func (e ErrRuleTokenMatchFailure) Error() string {
	message := `The rule failed to match the provided token as next in lexeme sequence.`

	return fmt.Sprintf(`[ErrRuleTokenMatchFailure] %s`, message)
}

type ErrRuleReferenceNotFound struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleReferenceNotFound(rule RuleKind, target RuleIdentity) error {
	return &ErrRuleReferenceNotFound{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: target,
	}
}

func (e ErrRuleReferenceNotFound) Error() string {
	message := `Rule reference is invalid and cannot be found against grammar.`

	return fmt.Sprintf(`[ErrRuleTokenMatchFailure] %s`, message)
}

type ErrRuleConcatenationEmptyRuleCollection struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleConcatenationEmptyRuleCollection(rule RuleKind) error {
	return &ErrRuleConcatenationEmptyRuleCollection{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: rule.GetIdentity(),
	}
}

func (e ErrRuleConcatenationEmptyRuleCollection) Error() string {
	message := `No choices are available, ruleset given is empty.`

	return fmt.Sprintf(`[ErrRuleConcatenationEmptyRuleCollection] %s`, message)
}

type ErrRuleChoiceEmptyRuleCollection struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleChoiceEmptyRuleCollection(rule RuleKind) error {
	return &ErrRuleChoiceEmptyRuleCollection{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: rule.GetIdentity(),
	}
}

func (e ErrRuleChoiceEmptyRuleCollection) Error() string {
	message := `No choices are available, ruleset given is empty.`

	return fmt.Sprintf(`[ErrRuleChoiceEmptyRuleCollection] %s`, message)
}

type ErrRuleChoiceNoneMatched struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleChoiceNoneMatched(rule RuleKind) error {
	return &ErrRuleChoiceNoneMatched{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: rule.GetIdentity(),
	}
}

func (e ErrRuleChoiceNoneMatched) Error() string {
	message := `No choices were matched.`

	return fmt.Sprintf(`[ErrRuleTokenMatchFailure] %s`, message)
}

type ErrRuleRepetitionMaximumZero struct {
	RuleType     string
	RuleIdentity RuleIdentity
}

func NewErrRuleRepetitionMaximumZero(rule RuleKind) error {
	return &ErrRuleRepetitionMaximumZero{
		RuleType:     fmt.Sprintf(`%T`, rule),
		RuleIdentity: rule.GetIdentity(),
	}
}

func (e ErrRuleRepetitionMaximumZero) Error() string {
	message := `Reptition rule cannot function with a zero maximum.`

	return fmt.Sprintf(`[ErrRuleRepetitionMaximumZero] %s`, message)
}
