package ast

type NodeSequence []NodeKind

type NodeKind interface {
	GetRule() RuleKind
	GetLexemes() LexemeSequence
	CountLexemes() int
	GetNodes() NodeSequence
	CountNodes() int
	IsEndNode() bool
}

type Node struct {
	Rule    RuleKind
	Lexemes LexemeSequence
	Nodes   NodeSequence
}

func (n Node) GetRule() RuleKind {
	return n.Rule
}

func (n Node) GetLexemes() LexemeSequence {
	return n.Lexemes
}

func (n Node) CountLexemes() int {
	return len(n.Lexemes)
}

func (n Node) GetNodes() NodeSequence {
	return n.Nodes
}

func (n Node) CountNodes() int {
	return len(n.Nodes)
}

func (n Node) IsEndNode() bool {
	return n.CountNodes() == 0 && n.CountLexemes() > 0
}
