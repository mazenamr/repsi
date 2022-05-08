package helpers

import (
	"fmt"
	"log"
	"math/big"
	"repsi/consts"
	"repsi/machines/nfa"
	"strconv"
)

func EpsilonClosure(s *nfa.State) []*nfa.State {
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

func Tokens(s []*nfa.State) []string {
	set := make(map[string]bool)
	for _, t := range s {
		for _, m := range t.Moves {
			if m.Token != consts.EmptyToken {
				if len(m.Token) > 1 {
					if m.Token[0] == '[' && m.Token[len(m.Token)-1] == ']' {
						if m.Token[1] != '^' {
							chars := ExpandCharset(m.Token[1 : len(m.Token)-1])
							for _, c := range chars {
								set[c] = true
							}
						} else {
							log.Fatal("[^] not supported yet")
						}
					} else {
						set[m.Token] = true
					}
				} else {
					set[m.Token] = true
				}
			}
		}
	}
	tokens := make([]string, 0, len(set))
	for t := range set {
		tokens = append(tokens, t)
	}
	return tokens
}

func Merge(s1 []*nfa.State, s2 []*nfa.State) []*nfa.State {
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

func Terminating(s []*nfa.State) bool {
	for _, t := range s {
		if t.Terminating {
			return true
		}
	}
	return false
}

func PrimeName(s []*nfa.State) string {
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
