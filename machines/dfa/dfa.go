package dfa

import (
	"fmt"
	"sort"
)

var count int64 = 0

type Machine struct {
	Name        string
	Moves       map[string]*Machine
	Terminating bool
}

func (m Machine) String() (r string) {
	machines := m.Machines()
	for _, s := range machines {
		if s.Terminating {
			r += ">>"
		}
		r += s.Name + ":\n"
		for token, to := range s.Moves {
			r += "\t" +
				fmt.Sprintf("-(%q)-> %s", token, to.Name) + "\n"
		}
	}
	return
}

func (m *Machine) Machines() []*Machine {
	set := make(map[*Machine]bool)
	set[m] = true

	change := true
	for change {
		change = false
		for s := range set {
			if set[s] {
				for _, t := range s.Moves {
					if _, ok := set[t]; !ok {
						set[t] = true
						change = true
					}
				}
				set[s] = false
			}
		}
	}

	states := make([]*Machine, 0, len(set))
	for s := range set {
		states = append(states, s)
	}

	sort.SliceStable(states, func(i, j int) bool {
		return states[i].Name < states[j].Name
	})

	return states
}

func (m *Machine) Tokens() []string {
	set := make(map[string]bool)
	machines := m.Machines()
	for _, n := range machines {
		for t := range n.Moves {
			set[t] = true
		}
	}
	tokens := make([]string, 0, len(set))
	for t := range set {
		tokens = append(tokens, t)
	}
	return tokens
}

func (m *Machine) Renumber() {
	count = 0
	for _, s := range m.Machines() {
		s.Name = fmt.Sprintf("S%d", count)
		count++
	}
}
