package ast

type LexemeSequence []LexemeKind

func (sequence LexemeSequence) Count() int {
	return len(sequence)
}

type LexemeKind interface {
	GetTokenIdentity() TokenIdentity
	GetTokenOffset() TokenOffset
	GetValue() string
	IsTokenIdentity(id TokenIdentity) bool
}

type Lexeme struct {
	Token  TokenIdentity
	Offset TokenOffset
	Value  string
}

func (l Lexeme) GetTokenIdentity() TokenIdentity {
	return l.Token
}

func (l Lexeme) GetTokenOffset() TokenOffset {
	return l.Offset
}

func (l Lexeme) GetValue() string {
	return l.Value
}

func (l Lexeme) IsTokenIdentity(id TokenIdentity) bool {
	return l.Token == id
}
