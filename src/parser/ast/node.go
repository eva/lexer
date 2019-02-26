package ast

type NodeType int

const (
	NodeTypeLexeme NodeType = 1
	NodeTypeRule   NodeType = 2
)

type NodeKind interface {
	GetNodeType() NodeType
	IsValid() bool
}
