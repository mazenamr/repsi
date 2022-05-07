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
	start := &State{Name: fmt.Sprintf("S%d", -1*count)}
	end := &State{Name: fmt.Sprintf("S%d", count+1), Terminating: true}
	count += 2
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: m.Start})
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: end})
	m.End.Moves = append(m.End.Moves, &Move{Token: consts.EmptyToken, To: end})
	m.End.Terminating = false
	return &Machine{start, end}
}

func (m *Machine) Plus() *Machine {
	start := &State{Name: fmt.Sprintf("S%d", -1*count)}
	end := &State{Name: fmt.Sprintf("S%d", count+1), Terminating: true}
	count += 2
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: m.Start})
	m.End.Moves = append(m.End.Moves, &Move{Token: consts.EmptyToken, To: m.Start})
	m.End.Moves = append(m.End.Moves, &Move{Token: consts.EmptyToken, To: end})
	m.End.Terminating = false
	return &Machine{start, end}
}

func (m *Machine) Star() *Machine {
	start := &State{Name: fmt.Sprintf("S%d", -1*count)}
	end := &State{Name: fmt.Sprintf("S%d", count+1), Terminating: true}
	count += 2
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: m.Start})
	start.Moves = append(start.Moves, &Move{Token: consts.EmptyToken, To: end})
	m.End.Moves = append(m.End.Moves, &Move{Token: consts.EmptyToken, To: m.Start})
	m.End.Moves = append(m.End.Moves, &Move{Token: consts.EmptyToken, To: end})
	m.End.Terminating = false
	return &Machine{start, end}
}

func Concat(m1, m2 *Machine) *Machine {
	m1.End.Moves = append(m1.End.Moves, &Move{Token: consts.EmptyToken, To: m2.Start})
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
