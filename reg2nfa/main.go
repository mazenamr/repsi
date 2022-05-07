package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <regex>\n", os.Args[0])
		os.Exit(1)
	}
	m := Parse(os.Args[1])
	fmt.Println(m)
	a := m.Abstract()
	a.Out("nfa")
}
