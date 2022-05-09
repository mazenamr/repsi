package main

import (
	"fmt"
	"os"
	"repsi/machines/abstract"
	"repsi/machines/dfa"
	"repsi/machines/nfa"
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

	a := abstract.Load("nfa.json")
	n := nfa.FromAbstract(a)

	d := dfa.Generate(n)
	d.Abstract().Out("dfa")

	d = d.Minimize()
	d.Abstract().Out("dfa-minimized")
}
