package dfa

import (
	"fmt"
	"math/big"
	"repsi/consts"
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

func (m *Machine) Minimize() *Machine {
	machines := m.Machines()
	machineGroup := make(map[*Machine]int)
	groups := make([][]*Machine, 0, len(machines))

	last := -1
	terminating, nonterminating := -1, -1
	for _, m := range machines {
		if !m.Terminating {
			if nonterminating == -1 {
				last++
				nonterminating = last
				groups = append(groups, make([]*Machine, 0, len(machines)))
			}
			machineGroup[m] = 0
			groups[nonterminating] = append(groups[nonterminating], m)
		} else {
			if terminating == -1 {
				last++
				terminating = last
				groups = append(groups, make([]*Machine, 0, len(machines)))
			}
			machineGroup[m] = 1
			groups[terminating] = append(groups[terminating], m)
		}
	}

	tokens := m.Tokens()
	primes := consts.Primes(len(tokens) * len(machines))

	vals := make(map[string]map[int]int)
	for i, t := range tokens {
		vals[t] = make(map[int]int)
		for j := 0; j < len(machines); j++ {
			vals[t][j] = primes[i*len(machines)+j]
		}
	}

	prev := -1
	for prev != last {
		prev = last
		for i, g := range groups {
			moveGroup := make(map[string]int)
			product := big.NewInt(1)
			for _, t := range tokens {
				if _, ok := g[0].Moves[t]; ok {
					product.Mul(product, big.NewInt(int64(vals[t][machineGroup[g[0].Moves[t]]])))
				}
			}
			moveGroup[fmt.Sprint(product)] = i
			for _, n := range g[1:] {
				product := big.NewInt(1)
				for _, t := range tokens {
					if _, ok := n.Moves[t]; ok {
						product.Mul(product, big.NewInt(int64(vals[t][machineGroup[n.Moves[t]]])))
					}
				}
				if _, ok := moveGroup[fmt.Sprint(product)]; !ok {
					last++
					machineGroup[n] = last
					moveGroup[fmt.Sprint(product)] = last
					groups = append(groups, make([]*Machine, 0, len(machines)))
				}
				if machineGroup[n] != i {
					groups[machineGroup[n]] = append(groups[machineGroup[n]], n)
					for j, g := range groups[i] {
						if g == n {
							groups[i] = append(groups[i][:j], groups[i][j+1:]...)
							break
						}
					}
				}
			}
		}
	}

	for _, g := range groups {
		for t := range g[0].Moves {
			g[0].Moves[t] = groups[machineGroup[g[0].Moves[t]]][0]
		}
	}

	n := groups[machineGroup[m]][0]
	n.Renumber()
	return n
}
