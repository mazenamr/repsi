package nfa

import "repsi/abstract"

func (m *Machine) Abstract() *abstract.AbstractMachine {
	a := &abstract.AbstractMachine{States: make(map[string]abstract.AbstractState)}
	a.StartingState = m.Start.Name
	for _, s := range m.States() {
		a.States[s.Name] = abstract.AbstractState{IsTerminatingState: s.Terminating, Moves: make(map[string][]string)}
		for _, t := range s.Moves {
			a.States[s.Name].AddMove(string(t.Token), t.To.Name)
		}
	}
	return a
}

func FromAbstract(a *abstract.AbstractMachine) *Machine {
	var end *State
	states := make(map[string]*State)
	for name, s := range a.States {
		states[name] = &State{
			name, make([]*Move, 0, len(s.Moves)),
			s.IsTerminatingState}

		if s.IsTerminatingState {
			end = states[name]
		}
	}

	for _, s := range states {
		for token, to := range a.States[s.Name].Moves {
			for _, t := range to {
				s.Moves = append(s.Moves, &Move{rune(token[0]), states[t]})
			}
		}
	}

	return &Machine{states[a.StartingState], end}
}