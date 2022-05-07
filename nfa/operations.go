package nfa

import (
	"fmt"
	"repsi/consts"
)

func (m *Machine) Concat(n *Machine) *Machine {
	m = Concat(m, n)
	return m
}

func (m *Machine) Optional() *Machine {
	return m
}

func (m *Machine) Plus() *Machine {
	return m
}

func (m *Machine) Star() *Machine {
	return m
}

func Concat(m1, m2 *Machine) *Machine {
	m1.End.Moves = append(m1.End.Moves, m2.Start.Moves...)
	m1.End.Terminating = false
	return &Machine{m1.Start, m2.End}
}

func Union(m1, m2 *Machine) *Machine {
	start := &State{Name: fmt.Sprintf("S%d", -1*count)}
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: m1.Start})
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: m2.Start})
	end := &State{Name: fmt.Sprintf("S%d", count+1), Terminating: true}
	count += 2
	m1.End.Moves = append(m1.End.Moves, &Move{Token: consts.EmptyToken, To: end})
	m1.End.Terminating = false
	m2.End.Moves = append(m2.End.Moves, &Move{Token: consts.EmptyToken, To: end})
	m2.End.Terminating = false
	return &Machine{start, end}
}
