package ast

type NodeSequence []NodeKind

func (sequence NodeSequence) Count() int {
	return len(sequence)
}

func (sequence NodeSequence) IsEmpty() bool {
	return sequence.Count() == 0
}

type NodeKind interface {
	GetRuleIdentity() RuleIdentity
	GetLexemeSequence() LexemeSequence
	CountLexemeSequence() int
	GetNodeSequence() NodeSequence
	CountNodeSequence() int
}

type Node struct {
	Rule    RuleIdentity
	Lexemes LexemeSequence
	Nodes   NodeSequence
}

func (n Node) GetRuleIdentity() RuleIdentity {
	return n.Rule
}

func (n Node) GetLexemeSequence() LexemeSequence {
	return n.Lexemes
}

func (n Node) CountLexemeSequence() int {
	return n.Lexemes.Count()
}

func (n Node) GetNodeSequence() NodeSequence {
	return n.Nodes
}

func (n Node) CountNodeSequence() int {
	return n.Nodes.Count()
}
