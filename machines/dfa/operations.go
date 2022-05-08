package dfa

import (
	"fmt"
	"log"
	"math/big"
	"repsi/consts"
	"repsi/machines/nfa"
	"strconv"
)

func Generate(m *nfa.Machine) *Machine {
	m.Prime()
	set := make(map[string]bool)
	machines := make(map[string]*Machine)
	closures := make(map[*Machine][]*nfa.State)

	startClosure := epsilonClosure(m.Start)
	start := &Machine{
		Name:        stateName(startClosure),
		Moves:       make(map[string]*Machine),
		Terminating: terminating(startClosure),
	}
	set[start.Name] = true
	machines[start.Name] = start
	closures[start] = startClosure

	change := true
	for change {
		change = false
		for s, m := range machines {
			if set[s] {
				for _, t := range tokens(closures[m]) {
					c := make([]*nfa.State, 0)
					states := make([]*nfa.State, 0)
					for _, state := range closures[m] {
						newStates := make([]*nfa.State, 0)
						for _, move := range state.Moves {
							if move.Token == t {
								newStates = append(newStates, move.To)
							}
						}
						states = merge(states, newStates)
					}
					for _, state := range states {
						c = merge(c, epsilonClosure(state))
					}
					name := stateName(c)
					if _, ok := set[name]; !ok {
						set[name] = true
						change = true
						n := &Machine{
							Name:        name,
							Moves:       make(map[string]*Machine),
							Terminating: terminating(c),
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

func merge(s1 []*nfa.State, s2 []*nfa.State) []*nfa.State {
	set := make(map[*nfa.State]bool)
	for _, s := range s1 {
		set[s] = true
	}
	for _, s := range s2 {
		set[s] = true
	}
	states := make([]*nfa.State, 0, len(set))
	for s := range set {
		states = append(states, s)
	}
	return states
}

func stateName(s []*nfa.State) string {
	product := big.NewInt(1)
	for _, t := range s {
		v, err := strconv.Atoi(t.Name[1:])
		if err != nil {
			log.Fatal(err)
		}
		product.Mul(product, big.NewInt(int64(v)))
	}
	return fmt.Sprintf("S%d", product)
}

func tokens(s []*nfa.State) []string {
	set := make(map[string]bool)
	for _, t := range s {
		for _, m := range t.Moves {
			if m.Token != consts.EmptyToken {
				set[m.Token] = true
			}
		}
	}
	tokens := make([]string, 0, len(set))
	for t := range set {
		tokens = append(tokens, t)
	}
	return tokens
}

func terminating(s []*nfa.State) bool {
	for _, t := range s {
		if t.Terminating {
			return true
		}
	}
	return false
}

func epsilonClosure(s *nfa.State) []*nfa.State {
	set := make(map[*nfa.State]bool)
	set[s] = true

	change := true
	for change {
		change = false
		for s := range set {
			if set[s] {
				for _, t := range s.Moves {
					if t.Token == consts.EmptyToken {
						if _, ok := set[t.To]; !ok {
							set[t.To] = true
							change = true
						}
					}
				}
				set[s] = false
			}
		}
	}

	states := make([]*nfa.State, 0, len(set))
	for s := range set {
		states = append(states, s)
	}

	return states
}
