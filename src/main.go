package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	examples "./examples/json"
	"./parser"
	"./parser/ast/printer"

	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Println(color.YellowString(fmt.Sprintf("Read %d characters:", len(text))))
	fmt.Println(fmt.Sprintf("%s", text))
	fmt.Println("")

	sequence, index, err := parser.Tokenise(examples.Grammar, text)

	fmt.Println(color.YellowString(fmt.Sprintf("Tokenised to %d tokens:", sequence.Count())))
	printer.PrintLexemeSequence(sequence)
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Index: %d", index))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}

	node, err := parser.ParseAnySequence(examples.Grammar, sequence)

	fmt.Println(color.YellowString("Node Tree:"))
	printer := printer.NodePrinter{}
	printer.Print(examples.Grammar, node)
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}
}
