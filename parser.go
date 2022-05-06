package main

import "repsi/nfa"

// turn a regex string into an nfa state machine

func Parse(s string) *nfa.Machine {
	m := nfa.EmptyMachine()
	for _, r := range s {
		m = m.Concat(nfa.TokenMachine(r))
	}
	return m
}
