package dfa

import (
	"repsi/machines/abstract"
	"repsi/machines/helpers"
)

func (m *Machine) Abstract() *abstract.Machine {
	a := &abstract.Machine{States: make(map[string]abstract.State)}
	a.StartingState = m.Name

	for _, s := range m.Machines() {
		a.States[s.Name] = abstract.State{IsTerminatingState: s.Terminating, Moves: make(map[string][]string)}

		tokens := make(map[*Machine][]string)

		for t, to := range s.Moves {
			if _, ok := tokens[to]; !ok {
				tokens[to] = []string{t}
			} else {
				tokens[to] = append(tokens[to], t)
			}
		}

		for to, ts := range tokens {
			a.States[s.Name].AddMove(helpers.GenerateCharacterSet(ts), to.Name)
		}
	}

	return a
}

func FromAbstract(a *abstract.Machine) *Machine {
	machines := make(map[string]*Machine)
	for name, s := range a.States {
		machines[name] = &Machine{
			name, make(map[string]*Machine),
			s.IsTerminatingState}
	}

	for _, s := range machines {
		for token, to := range a.States[s.Name].Moves {
			for _, t := range to {
				s.Moves[token] = machines[t]
			}
		}
	}

	return machines[a.StartingState]
}
