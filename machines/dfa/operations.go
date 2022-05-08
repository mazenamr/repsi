package dfa

import (
	"repsi/machines/helpers"
	"repsi/machines/nfa"
)

func Generate(m *nfa.Machine) *Machine {
	m.Prime()
	set := make(map[string]bool)
	machines := make(map[string]*Machine)
	closures := make(map[*Machine][]*nfa.State)

	startClosure := helpers.EpsilonClosure(m.Start)
	start := &Machine{
		Name:        helpers.PrimeName(startClosure),
		Moves:       make(map[string]*Machine),
		Terminating: helpers.Terminating(startClosure),
	}
	set[start.Name] = true
	machines[start.Name] = start
	closures[start] = startClosure

	change := true
	for change {
		change = false
		for s, m := range machines {
			if set[s] {
				for _, t := range helpers.Tokens(closures[m]) {
					c := make([]*nfa.State, 0)
					states := make([]*nfa.State, 0)
					for _, state := range closures[m] {
						newStates := make([]*nfa.State, 0)
						for _, move := range state.Moves {
							if helpers.Match(t, move.Token) {
								newStates = append(newStates, move.To)
							}
						}
						states = helpers.Merge(states, newStates)
					}
					for _, state := range states {
						c = helpers.Merge(c, helpers.EpsilonClosure(state))
					}
					name := helpers.PrimeName(c)
					if _, ok := set[name]; !ok {
						set[name] = true
						change = true
						n := &Machine{
							Name:        name,
							Moves:       make(map[string]*Machine),
							Terminating: helpers.Terminating(c),
						}
						set[n.Name] = true
						machines[n.Name] = n
						closures[n] = c
					}
					m.Moves[t] = machines[name]
				}
				set[s] = false
			}
		}
	}

	start.Renumber()
	return start
}
