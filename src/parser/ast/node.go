package ast

type NodeType int

const (
	NodeTypeNull   NodeType = 0
	NodeTypeLexeme NodeType = 1
	NodeTypeRule   NodeType = 2
)

type NodeKind interface {
	GetNodeType() NodeType
	IsValid() bool
}

type NodeNull struct{}

func (NodeNull) GetNodeType() NodeType {
	return NodeTypeNull
}

func (node NodeNull) IsValid() bool {
	return false
}
