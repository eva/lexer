package ast

import "strings"

type TokenLiteral struct {
	identity   TokenIdentity
	namespace  Namespace
	transition string
	literal    string
}

func NewTokenLiteral(identity TokenIdentity, namespace Namespace, literal string) TokenLiteral {
	token := TokenLiteral{}
	token.identity = identity
	token.namespace = namespace
	token.literal = literal

	return token
}

func (token TokenLiteral) Identity() TokenIdentity {
	return token.identity
}

func (token TokenLiteral) Namespace() Namespace {
	return token.namespace
}

func (token TokenLiteral) Transition() (change bool, namespace string) {
	namespace = token.transition

	if namespace == "" {
		change = false
	} else {
		change = true
	}

	return
}

func (token TokenLiteral) Literal() string {
	return token.literal
}

func (token TokenLiteral) Match(input string) (bool, TokenOffset) {
	index := strings.Index(input, token.literal)

	if index == -1 {
		return false, NoTokenOffset
	}

	length := len(token.literal)
	offset := TokenOffset{index, length}

	return true, offset
}
