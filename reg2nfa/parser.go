package main

import (
	"repsi/consts"
	"repsi/nfa"
)

func Parse(s string) *nfa.Machine {
	m := nfa.EmptyMachine()
	for _, r := range s {
		m = m.Concat(nfa.TokenMachine(r))
	}
	return m
}

func Preprocess(s string) []*Token {
	tokens := make([]*Token, 0, len(s))
	return tokens
}
