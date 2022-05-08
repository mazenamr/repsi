package main

import (
	"fmt"
	"os"
	"repsi/machines/abstract"
	"repsi/machines/dfa"
	"repsi/machines/nfa"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s convert <input-file> <output-file>\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[1] != "convert" {
		fmt.Fprintf(os.Stderr, "Usage: %s convert <input-file> <output-file>\n", os.Args[0])
		os.Exit(1)
	}

	a := abstract.Load(os.Args[2])
	n := nfa.FromAbstract(a)
	d := dfa.Generate(n)
	d.Abstract().Out(os.Args[3])
	d = d.Minimize()
	d.Abstract().Out(os.Args[3] + "-min")
}
