package nfa

import (
	"fmt"
	"repsi/consts"
)

func EmptyMachine() *Machine {
	end := &State{Name: fmt.Sprintf("S%d", count+1), Terminating: true}
	start := &State{Name: fmt.Sprintf("S%d", count), Moves: []*Move{{Token: consts.EmptyToken, To: end}}}
	count += 2
	return &Machine{start, end}
}

func TokenMachine(t rune) *Machine {
	m := EmptyMachine()
	m.Start.Moves = []*Move{{Token: t, To: m.End}}
	return m
}

func PlusMachine() *Machine {
	m := EmptyMachine()
	return m
}

func StarMachine() *Machine {
	m := EmptyMachine()
	return m
}
