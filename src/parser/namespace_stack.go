package parser

import "./ast"

// A NamespaceStack represents a history of namespace traversal to get this token.
// When a token shifts a namespace the stack will shrink.
type NamespaceStack struct {
	stack []ast.NamespaceKind
}

func (ns NamespaceStack) Count() int {
	return len(ns.stack)
}

func (ns NamespaceStack) Current() ast.NamespaceKind {
	count := ns.Count()

	if count == 0 {
		return nil
	}

	return ns.stack[count-1]
}

func (ns *NamespaceStack) Shift() {
	count := ns.Count()

	ns.stack = ns.stack[:count-1]
}

func (ns *NamespaceStack) Register(namespace ast.NamespaceKind) {
	ns.stack = append(ns.stack, namespace)
}
