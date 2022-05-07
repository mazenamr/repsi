package dfa

import (
	"fmt"
	"sort"
)

var count int = 0

type Machine struct {
	Name        string
	Moves       map[string]*Machine
	Terminating bool
}

func (m Machine) String() (r string) {
	states := m.States()
	for _, s := range states {
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

func (m *Machine) States() []*Machine {
	set := make(map[*Machine]bool)
	set[m] = true

	for {
		change := false
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
		if !change {
			break
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

func (m *Machine) Renumber() {
	count = 0
	for _, s := range m.States() {
		s.Name = fmt.Sprintf("S%d", count)
		count++
	}
}
