package main

import (
	"fmt"
	"os"
	"repsi/parser"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s convert <regex>\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[1] != "convert" {
		fmt.Fprintf(os.Stderr, "Usage: %s convert <regex>\n", os.Args[0])
		os.Exit(1)
	}

	m := parser.Parse(os.Args[2])
	m.Abstract().Out("nfa")
}
