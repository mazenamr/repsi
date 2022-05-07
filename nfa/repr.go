package nfa

import (
	"fmt"
)

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
