package ast

type NodeRuleKind interface {
	GetRuleIdentity() RuleIdentity
	GetNodeSequence() NodeSequence
	IsEmpty() bool
	Count() int
}

type NodeRule struct {
	Rule  RuleIdentity
	Nodes NodeSequence
}

func (NodeRule) GetNodeType() NodeType {
	return NodeTypeRule
}

func (n NodeRule) IsValid() bool {
	return n.IsEmpty() == false
}

func (n NodeRule) GetRuleIdentity() RuleIdentity {
	return n.Rule
}

func (n NodeRule) GetNodeSequence() NodeSequence {
	return n.Nodes
}

func (n NodeRule) IsEmpty() bool {
	return n.Nodes.Count() > 0
}

func (n NodeRule) Count() int {
	return n.Nodes.Count()
}
