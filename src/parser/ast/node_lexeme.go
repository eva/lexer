package ast

type NodeLexemeKind interface {
	GetTokenIdentity() TokenIdentity
	IsTokenIdentity(id TokenIdentity) bool
	GetTokenOffset() TokenOffset
	GetValue() string
}

type NodeLexeme struct {
	Token  TokenIdentity
	Offset TokenOffset
	Value  string
}

func (n NodeLexeme) GetNodeType() NodeType {
	return NodeTypeLexeme
}

func (n NodeLexeme) IsValid() bool {
	return true
}

func (n NodeLexeme) GetTokenIdentity() TokenIdentity {
	return n.Token
}

func (n NodeLexeme) IsTokenIdentity(id TokenIdentity) bool {
	return n.Token == id
}

func (n NodeLexeme) GetTokenOffset() TokenOffset {
	return n.Offset
}

func (n NodeLexeme) GetValue() string {
	return n.Value
}
