package main

import (
	"fmt"
	"repsi/nfa"
)

func main() {
	m1 := Parse("hello, ")
	m2 := Parse("world")
	m3 := Parse("gopher")
	m4 := nfa.Union(m2, m3)
	m5 := Parse("!")
	m := m1.Concat(m4).Concat(m5)
	fmt.Printf("%s\n\n", m)
	a := m.Abstract()
	// a.States["S13"].Moves["x"] = append(make([]string, 0), "S45")
	a.States["S13"].AddMove("x", "S45")
	for c := 'A'; c <= 'Z'; c++ {
		a.States["S3"].AddMove(string(c), "S5")
	}
	a.Write("hello")
	a.Draw("hello")
}
