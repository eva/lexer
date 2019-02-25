package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	examples "./examples/json"
	"./parser"

	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	fmt.Println(color.YellowString("Read:"))
	fmt.Println(fmt.Sprintf("%#v", text))
	fmt.Println("")

	sequence, index, err := parser.Tokenise(examples.Grammar, text)

	value, _ := json.Marshal(sequence)
	fmt.Println(color.YellowString("Token Sequence:"))
	fmt.Println(string(value))
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Index: %d", index))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}

	node, err := parser.ParseAnySequence(examples.Grammar, sequence)

	value, _ = json.MarshalIndent(node, "", " ")
	fmt.Println(color.YellowString("Node:"))
	fmt.Println(string(value))
	fmt.Println("")

	if err != nil {
		fmt.Println(color.RedString("Error:"))
		fmt.Println(fmt.Sprintf("Message: %#v", err))
	}
}
