package ast

type NodeSequence []NodeKind

type NodeKind interface {
	GetRuleIdentity() RuleIdentity
	GetLexemes() LexemeSequence
	CountLexemes() int
	GetNodes() NodeSequence
	CountNodes() int
	IsEndNode() bool
}

type Node struct {
	Rule    RuleIdentity
	Lexemes LexemeSequence
	Nodes   NodeSequence
}

func (n Node) GetRuleIdentity() RuleIdentity {
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
