package ast

type LexemeKind interface {
	GetNamespaceIdentity() NamespaceIdentity
	GetTokenIdentity() TokenIdentity
	IsTokenIdentity(id TokenIdentity) bool
	GetTokenOffset() TokenOffset
	GetValue() string
	IsValid() bool
}

type Lexeme struct {
	Namespace NamespaceIdentity
	Token     TokenIdentity
	Offset    TokenOffset
	Value     string
}

func (lexeme Lexeme) GetNamespaceIdentity() NamespaceIdentity {
	return lexeme.Namespace
}

func (lexeme Lexeme) GetTokenIdentity() TokenIdentity {
	return lexeme.Token
}

func (lexeme Lexeme) IsTokenIdentity(id TokenIdentity) bool {
	return lexeme.Token == id
}

func (lexeme Lexeme) GetTokenOffset() TokenOffset {
	return lexeme.Offset
}

func (lexeme Lexeme) GetValue() string {
	return lexeme.Value
}

func (lexeme Lexeme) IsValid() bool {
	if lexeme.Token == TokenIdentityNone {
		return false
	}

	if lexeme.Offset.IsValid() == false {
		return false
	}

	if len(lexeme.Value) == 0 {
		return false
	}

	return true
}
