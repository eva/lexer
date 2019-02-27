package ast

// NodeRuleKind represents a node that is the result of a rule match that has multiple children.
// An example might be something like a concatenation rule, where many children nodes are required.
// Also grouping tokens by their rule allows for a better tree output.
type NodeRuleKind interface {
	GetRuleIdentity() RuleIdentity
	GetNodeSequence() NodeSequence
	IsEmpty() bool
	Count() int
}

// NodeRule is an implementation of NodeRuleKind, this node will likely contain no lexical information.
// All data available in the Lexeme should be available here as this is a matched token.
type NodeRule struct {
	Rule  RuleIdentity
	Nodes NodeSequence
}

// GetNodeType will simply return the NodeTypeRule node type.
func (NodeRule) GetNodeType() NodeType {
	return NodeTypeRule
}

// IsValid will check the values of the node and return a boolean indicating its validity.
// A node that is initialised with default values should always be considered invalid.
func (node NodeRule) IsValid() bool {
	if node.Rule == RuleIdentityNone {
		return false
	}

	if node.IsEmpty() {
		return false
	}

	return true
}

func (node NodeRule) GetRuleIdentity() RuleIdentity {
	return node.Rule
}

func (node NodeRule) GetNodeSequence() NodeSequence {
	return node.Nodes
}

func (node NodeRule) IsEmpty() bool {
	return node.Nodes.Count() == 0
}

func (node NodeRule) Count() int {
	return node.Nodes.Count()
}
