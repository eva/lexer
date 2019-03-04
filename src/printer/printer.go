package printer

import (
	"fmt"
	"strings"

	"../parser/ast"
	"github.com/fatih/color"
)

func PrintLexemeSequence(grammar ast.GrammarKind, sequence ast.LexemeSequence) {
	for _, lexeme := range sequence {
		fmt.Println(prepareLexemeKindOutput(grammar, lexeme))
	}
}

type NodePrinter struct {
	indent int
}

func (printer NodePrinter) Print(grammar ast.GrammarKind, node ast.NodeKind) {
	if rule, ok := node.(ast.NodeRule); ok {
		printer.printNodeRule(grammar, rule)
		return
	}

	if lexeme, ok := node.(ast.NodeLexeme); ok {
		printer.printNodeLexeme(grammar, lexeme)
		return
	}
}

func (printer NodePrinter) PrintSequence(grammar ast.GrammarKind, sequence ast.NodeSequence) {
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

func (printer NodePrinter) printNodeLexeme(grammar ast.GrammarKind, node ast.NodeLexeme) {
	indent := printer.getIndent()
	fmt.Println(fmt.Sprintf(`%s%s`, indent, prepareLexemeKindOutput(grammar, node)))
}

func (printer NodePrinter) printNodeRule(grammar ast.GrammarKind, node ast.NodeRule) {
	fmt.Println(fmt.Sprintf(`%s%s`, printer.getIndent(), prepareNodeRuleOutput(grammar, node)))
	printer.incrementIndent()
	printer.PrintSequence(grammar, node.GetNodeSequence())
	printer.decrementIndent()
}

func prepareNodeRuleOutput(grammar ast.GrammarKind, node ast.NodeRule) string {
	rid := node.GetRuleIdentity()

	_, rule := grammar.FindRule(rid)
	name := rule.GetName()

	if name == "" {
		name = "unknown"
	}

	return fmt.Sprintf(
		`%s: %s/%s`,
		color.WhiteString(`%T`, node),
		color.YellowString(`%s`, name),
		fmt.Sprintf(`%d`, rid),
	)
}

func prepareLexemeKindOutput(grammar ast.GrammarKind, lexeme ast.LexemeKind) string {
	nsid := lexeme.GetNamespaceIdentity()
	tid := lexeme.GetTokenIdentity()
	offset := lexeme.GetTokenOffset()
	value := lexeme.GetValue()

	_, namespace := grammar.FindNamespace(nsid)
	_, token := namespace.FindToken(tid)
	name := token.GetName()

	if name == "" {
		name = "unknown"
	}

	return fmt.Sprintf(
		`%s: %s/%s(%s) [%s:%s] %s`,
		color.WhiteString(`%T`, lexeme),
		color.MagentaString(`@%s`, nsid),
		color.YellowString(`%s`, name),
		fmt.Sprintf(`%d`, tid),
		color.GreenString(`%d`, offset[0]),
		color.GreenString(`%d`, offset[1]),
		color.CyanString(`%#v`, value),
	)
}
