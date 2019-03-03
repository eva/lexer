package printer

import (
	"fmt"
	"strings"

	ast ".."
)

func PrintLexemeSequence(sequence ast.LexemeSequence) {
	for _, lexeme := range sequence {
		nsid := lexeme.GetNamespaceIdentity()

		tid := lexeme.GetTokenIdentity()

		offset := lexeme.GetTokenOffset()
		value := lexeme.GetValue()

		fmt.Println(fmt.Sprintf(`%T: @%s [t:%d] [%d:%d] %#v`, lexeme, nsid, tid, offset[0], offset[1], value))
	}
}

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
	indent := printer.getIndent()

	nsid := node.GetNamespaceIdentity()

	tid := node.GetTokenIdentity()

	offset := node.GetTokenOffset()
	value := node.GetValue()

	fmt.Println(fmt.Sprintf(`%s%T: @%s [t:%d] [%d:%d] %#v`, indent, node, nsid, tid, offset[0], offset[1], value))
}

func (printer NodePrinter) printNodeRule(grammar ast.Grammar, node ast.NodeRule) {
	fmt.Println(fmt.Sprintf(`%s%T: [r:%d]`, printer.getIndent(), node, node.GetRuleIdentity()))
	printer.incrementIndent()
	printer.PrintSequence(grammar, node.GetNodeSequence())
	printer.decrementIndent()
}
