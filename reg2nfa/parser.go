package main

import (
	"repsi/consts"
	"repsi/nfa"
)

func Parse(s string) *nfa.Machine {
	m := nfa.EmptyMachine()
	for _, r := range s {
		m = m.Concat(nfa.TokenMachine(r))
	}
	return m
}

func Preprocess(s string) []*Token {
	tokens := make([]*Token, 0, len(s))
	i := 0
	for i < len(s) {
		if !Special[s[i]] {
			switch s[i] {
			case '\\':
				i++
				tokens = append(tokens, &Token{Value: string(s[i]), Operation: Literal})
			case '.':
				tokens = append(tokens, &Token{Value: string(consts.WildcardToken), Operation: Wildcard})
			case '[':
				token := "["
				for i < len(s) && s[i] != ']' {
					i++
					token += string(s[i])
				}
				if token[len(token)-1] != ']' {
					panic("invalid regex")
				}
				tokens = append(tokens, &Token{Value: token, Operation: CharSet})
			case '{':
				min := ""
				max := ""
				i++
				for i < len(s) && s[i] != ',' && s[i] != '}' {
					min += string(s[i])
					i++
				}
				if i < len(s) && s[i] == ',' {
					i++
					for i < len(s) && s[i] != '}' {
						max += string(s[i])
						i++
					}
				}
				if min == "" {
					min = "0"
				}

			}

			if i+1 < len(s) && !Operation[s[i+1]] {
				tokens = append(tokens, &Token{Operation: Concat})
				i++
			}
		}
	}
	return tokens
}
