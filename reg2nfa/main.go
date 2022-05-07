package main

import (
	"fmt"
	"os"
	"repsi/reg2nfa/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <regex>\n", os.Args[0])
		os.Exit(1)
	}
	m := parser.Parse(os.Args[1])
	fmt.Println(m)
	a := m.Abstract()
	a.Out("nfa")
}
