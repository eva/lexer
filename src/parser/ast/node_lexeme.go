package ast

// NodeLexemeKind represents a node that is crossed with a lexeme.
// Simply this node is a matched token and should mirror the lexeme it was created from.
type NodeLexemeKind interface {
	GetNamespaceIdentity() NamespaceIdentity
	GetTokenIdentity() TokenIdentity
	IsTokenIdentity(id TokenIdentity) bool
	GetTokenOffset() TokenOffset
	GetValue() string
}

func NewLexemeNode(lexeme LexemeKind) NodeKind {
	return NodeLexeme{
		Namespace: lexeme.GetNamespaceIdentity(),
		Token:     lexeme.GetTokenIdentity(),
		Offset:    lexeme.GetTokenOffset(),
		Value:     lexeme.GetValue(),
	}
}

// NodeLexeme is an instance of NodeLexemeKind implementing the same interface as LexemeKind almost.
// All data available in the Lexeme should be available here as this is a matched token.
type NodeLexeme struct {
	Namespace NamespaceIdentity
	Token     TokenIdentity
	Offset    TokenOffset
	Value     string
}

// GetNodeType will simply return the NodeTypeLexeme node type.
func (NodeLexeme) GetNodeType() NodeType {
	return NodeTypeLexeme
}

// IsValid will check the values of the node and return a boolean indicating its validity.
// A node that is initialised with default values should always be considered invalid.
func (node NodeLexeme) IsValid() bool {
	if node.Token == TokenIdentityNone {
		return false
	}

	if node.Offset.IsValid() == false {
		return false
	}

	if len(node.Value) == 0 {
		return false
	}

	return true
}

func (node NodeLexeme) GetNamespaceIdentity() NamespaceIdentity {
	return node.Namespace
}

// GetTokenIdentity will return the token identity matched.
func (node NodeLexeme) GetTokenIdentity() TokenIdentity {
	return node.Token
}

func (lexeme NodeLexeme) IsTokenIdentity(id TokenIdentity) bool {
	return lexeme.Token == id
}

// GetTokenOffset will return the offset data captured when the token was matched.
func (node NodeLexeme) GetTokenOffset() TokenOffset {
	return node.Offset
}

// GetValue will return the value of the token that was matched.
func (node NodeLexeme) GetValue() string {
	return node.Value
}
