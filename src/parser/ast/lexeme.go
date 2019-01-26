package ast

type LexemeSequence []Lexeme

type Lexeme struct {
	Token  TokenKind
	Offset TokenOffset
	Value  string
}

func (l Lexeme) IsToken(id TokenIdentity) bool {
	tokenid := l.Token.GetIdentity()

	return tokenid == id
}
