package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	examples "./examples/json"
	"./parser"
	"./printer"

	"github.com/fatih/color"
)

func main() {
	grammar := examples.Grammar

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Println(color.YellowString(fmt.Sprintf("Read %d characters:", len(text))))
	fmt.Println(fmt.Sprintf("%s", text))
	fmt.Println("")

	tokenisestart := time.Now()
	sequence, index, err := parser.Tokenise(grammar, text)
	tokeniseduration := time.Since(tokenisestart)

	fmt.Println(color.YellowString(fmt.Sprintf("Tokenised to %d tokens: (time: %s)", sequence.Count(), tokeniseduration)))
	printer.PrintLexemeSequence(grammar, sequence)
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Index: %d", index))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}

	parsestart := time.Now()
	node, err := parser.ParseAnySequence(grammar, sequence)
	parseduration := time.Since(parsestart)

	fmt.Println(color.YellowString(fmt.Sprintf("Node Tree: (time: %s)", parseduration)))
	printer := printer.NodePrinter{}
	printer.Print(grammar, node)
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}
}
