package ast

// LexemeSequence is collection of lexeme kind.
type LexemeSequence []LexemeKind

func (sequence LexemeSequence) Count() int {
	return len(sequence)
}

func (sequence LexemeSequence) IsEmpty() bool {
	return len(sequence) == 0
}

func (sequence LexemeSequence) IsValid() bool {
	for _, lexeme := range sequence {
		if lexeme.IsValid() == false {
			return false
		}
	}

	return true
}

type LexemeKind interface {
	GetTokenIdentity() TokenIdentity
	IsTokenIdentity(id TokenIdentity) bool
	GetTokenOffset() TokenOffset
	GetValue() string
	IsValid() bool
}

type Lexeme struct {
	Token  TokenIdentity
	Offset TokenOffset
	Value  string
}

func (l Lexeme) GetTokenIdentity() TokenIdentity {
	return l.Token
}

func (l Lexeme) IsTokenIdentity(id TokenIdentity) bool {
	return l.Token == id
}

func (l Lexeme) GetTokenOffset() TokenOffset {
	return l.Offset
}

func (l Lexeme) GetValue() string {
	return l.Value
}

func (lexeme Lexeme) IsValid() bool {
	if lexeme.Token == InvalidTokenIdentity {
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
