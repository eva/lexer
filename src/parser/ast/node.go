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

type NodeNull struct {
	Nodes NodeSequence
}

func (NodeNull) GetNodeType() NodeType {
	return NodeTypeNull
}

func (node NodeNull) IsValid() bool {
	if node.IsEmpty() {
		return false
	}

	return true
}

func (node NodeNull) GetNodeSequence() NodeSequence {
	return node.Nodes
}

func (node NodeNull) IsEmpty() bool {
	return node.Nodes.Count() == 0
}

func (node NodeNull) Count() int {
	return node.Nodes.Count()
}
