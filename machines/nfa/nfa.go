package nfa

import (
	"fmt"
	"log"
	"repsi/consts"
	"sort"
	"strconv"
)

var count int64 = 0

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

func (m Machine) String() (r string) {
	states := m.States()
	for _, s := range states {
		r += s.String()
	}
	return
}

func (s State) String() (r string) {
	if s.Terminating {
		r += ">>"
	}
	r += s.Name + ":\n"
	for _, t := range s.Moves {
		r += "\t" + t.String() + "\n"
	}
	return
}

func (t Move) String() string {
	return fmt.Sprintf("-(%q)-> %s", t.Token, t.To.Name)
}

func (m *Machine) States() []*State {
	set := make(map[*State]bool)
	set[m.Start] = true

	change := true
	for change {
		change = false
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
	}

	states := make([]*State, 0, len(set))
	for s := range set {
		states = append(states, s)
	}

	sort.SliceStable(states, func(i, j int) bool {
		a, e1 := strconv.Atoi(states[i].Name[1:])
		b, e2 := strconv.Atoi(states[j].Name[1:])
		if e1 != nil || e2 != nil {
			log.Fatal("invalid state name")
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

func (m *Machine) Prime() {
	m.Renumber()
	primes := consts.Primes(len(m.States()))
	count = primes[len(primes)-1]
	states := m.States()
	for i, s := range states {
		s.Name = fmt.Sprintf("S%d", primes[i])
	}
}
