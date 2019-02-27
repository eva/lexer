package printer

import (
	"fmt"
	"strings"

	".."
)

type NodePrinter struct {
	indent int
}

func (printer NodePrinter) Print(grammar ast.Grammar, node ast.NodeKind) {
	if rule, ok := node.(ast.NodeRule); ok {
		printer.printNodeRule(grammar, rule)
		return
	}

	if lexeme, ok := node.(ast.NodeLexeme); ok {
		printer.printNodeLexeme(grammar, lexeme)
		return
	}
}

func (printer NodePrinter) PrintSequence(grammar ast.Grammar, sequence ast.NodeSequence) {
	for _, node := range sequence {
		printer.Print(grammar, node)
	}
}

func (printer *NodePrinter) incrementIndent() {
	printer.indent = printer.indent + 1
}

func (printer *NodePrinter) decrementIndent() {
	printer.indent = printer.indent - 1
}

func (printer NodePrinter) getIndent() string {
	return strings.Repeat(` `, printer.indent)
}

func (printer NodePrinter) printNodeLexeme(grammar ast.Grammar, node ast.NodeLexeme) {
	fmt.Println(fmt.Sprintf(`%s%T: [Token: %d] %s`, printer.getIndent(), node, node.GetTokenIdentity(), node.GetValue()))
}

func (printer NodePrinter) printNodeRule(grammar ast.Grammar, node ast.NodeRule) {
	fmt.Println(fmt.Sprintf(`%s%T: [Rule: %d]`, printer.getIndent(), node, node.GetRuleIdentity()))
	printer.incrementIndent()
	printer.PrintSequence(grammar, node.GetNodeSequence())
	printer.decrementIndent()
}
