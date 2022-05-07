package nfa

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	count int = 0
)

type Machine struct {
	Start *State
	End   *State
}

type State struct {
	Name        string
	Moves       []*Move
	Terminating bool
}

type Move struct {
	Token string
	To    *State
}

func (m *Machine) States() []*State {
	set := make(map[*State]bool)
	set[m.Start] = true

	for {
		change := false
		for s := range set {
			if set[s] {
				for _, t := range s.Moves {
					if _, ok := set[t.To]; !ok {
						set[t.To] = true
						change = true
					}
				}
				set[s] = false
			}
		}
		if !change {
			break
		}
	}

	states := make([]*State, 0, len(set))
	for s := range set {
		states = append(states, s)
	}

	sort.SliceStable(states, func(i, j int) bool {
		a, e1 := strconv.Atoi(states[i].Name[1:])
		b, e2 := strconv.Atoi(states[j].Name[1:])
		if e1 != nil || e2 != nil {
			panic("invalid state name")
		}
		return a < b
	})

	return states
}

func (m *Machine) Copy() *Machine {
	equivalent := make(map[*State]*State)
	for _, s := range m.States() {
		equivalent[s] = &State{Name: s.Name, Moves: make([]*Move, 0, len(s.Moves))}
	}
	for _, s := range m.States() {
		for _, m := range s.Moves {
			equivalent[s].Moves = append(equivalent[s].Moves, &Move{Token: m.Token, To: equivalent[m.To]})
		}
	}
	return &Machine{equivalent[m.Start], equivalent[m.End]}
}

func (m *Machine) Renumber() {
	count = 0
	for _, s := range m.States() {
		s.Name = fmt.Sprintf("S%d", count)
		count++
	}
}
