package parser

import (
	"fmt"
	"log"
	"repsi/consts"
	"repsi/machines/nfa"
	"strconv"
)

func Parse(s string) *nfa.Machine {
	if len(s) == 0 {
		return nfa.EmptyMachine()
	}
	tokens := Preprocess(s)
	if !CheckBrakets(tokens) {
		log.Fatal("invalid regex")
	}
	queue := Postfix(tokens)
	stack := make([]*nfa.Machine, 0, len(queue))
	for _, token := range queue {
		switch token.Type {
		case Literal:
			stack = append(stack, nfa.TokenMachine(token.Value))
		case Wildcard:
			stack = append(stack, nfa.TokenMachine(consts.WildcardToken))
		case CharSet:
			stack = append(stack, nfa.TokenMachine(token.Value))
		case Union:
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, nfa.Union(l, r))
		case Concat:
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, nfa.Concat(l, r))
		case Repeat:
			count, err := strconv.Atoi(token.Value)
			if err != nil {
				log.Fatal("invalid token")
			}
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			c := r.Copy()
			for i := 0; i < count; i++ {
				r = r.Concat(c.Copy())
			}
			stack = append(stack, r)
		case Optional:
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, r.Optional())
		case Plus:
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, r.Plus())
		case Star:
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, r.Star())
		default:
			log.Fatal("invalid token")
		}
	}
	if len(stack) != 1 {
		log.Fatal("invalid regex")
	}
	stack[0].Renumber()
	return stack[0]
}

func Preprocess(s string) []*Token {
	tokens := make([]*Token, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			tokens = append(tokens, &Token{Value: consts.WildcardToken, Type: Wildcard})
		case '*':
			tokens = append(tokens, &Token{Type: Star})
		case '+':
			tokens = append(tokens, &Token{Type: Plus})
		case '?':
			tokens = append(tokens, &Token{Type: Optional})
		case '|':
			tokens = append(tokens, &Token{Type: Union})
			if i+1 < len(s) && Operator[CharType[s[i+1]]] {
				log.Fatal("invalid regex")
			}
		case '(':
			tokens = append(tokens, &Token{Type: OpenGroup})
		case ')':
			tokens = append(tokens, &Token{Type: CloseGroup})
		case '[':
			token := "["
			for i < len(s) && s[i] != ']' {
				i++
				token += string(s[i])
			}
			if token[len(token)-1] != ']' {
				log.Fatal("invalid regex")
			}
			tokens = append(tokens, &Token{Value: token, Type: CharSet})
		case '{':
			min := ""
			i++
			for i < len(s) && s[i] != ',' && s[i] != '}' {
				min += string(s[i])
				i++
			}
			if i < len(s) && s[i] == ',' {
				max := ""
				i++
				for i < len(s) && s[i] != '}' {
					max += string(s[i])
					i++
				}
				if min == "" {
					min = "0"
				}
				minval, err := strconv.Atoi(min)
				if err != nil {
					log.Fatal("invalid regex")
				}
				tokens = append(tokens, &Token{Value: min, Type: Repeat})
				if max == "" {
					tokens = append(tokens, &Token{Type: Plus})
				} else {
					maxval, err := strconv.Atoi(max)
					if err != nil {
						log.Fatal("invalid regex")
					}
					if maxval < minval {
						log.Fatal("invalid regex")
					}
					log.Fatal("{min,max} not supported yet")
				}
			} else if i < len(s) && s[i] == '}' {
				minval, err := strconv.Atoi(min)
				if err != nil {
					log.Fatal("invalid regex")
				}
				if minval < 0 {
					log.Fatal("invalid regex")
				} else if minval == 0 {
					tokens = append(tokens, &Token{Type: Optional})
				} else {
					tokens = append(tokens, &Token{Value: fmt.Sprint(minval - 1), Type: Repeat})
				}
			} else {
				log.Fatal("invalid regex")
			}
		case '\\':
			i++
			tokens = append(tokens, &Token{Value: string(s[i]), Type: Literal})
		default:
			tokens = append(tokens, &Token{Value: string(s[i]), Type: Literal})
		}

		if CharType[s[i]] != Union && CharType[s[i]] != OpenGroup {
			if i+1 < len(s) && !Operator[CharType[s[i+1]]] && CharType[s[i+1]] != CloseGroup {
				tokens = append(tokens, &Token{Type: Concat})
			}
		}
	}
	return tokens
}

func CheckBrakets(tokens []*Token) bool {
	stack := make([]*Token, 0, len(tokens))
	for _, token := range tokens {
		switch token.Type {
		case OpenGroup:
			stack = append(stack, token)
		case CloseGroup:
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func Postfix(tokens []*Token) []*Token {
	stack := make([]*Token, 0, len(tokens))
	queue := make([]*Token, 0, len(tokens))
	for _, token := range tokens {
		if token.Type == Literal || token.Type == Wildcard || token.Type == CharSet {
			queue = append(queue, token)
		} else if Operator[token.Type] {
			for len(stack) > 0 {
				t := stack[len(stack)-1]
				if Precedence[t.Type] >= Precedence[token.Type] {
					stack = stack[:len(stack)-1]
					queue = append(queue, t)
				} else {
					break
				}
			}
			stack = append(stack, token)
		} else if token.Type == OpenGroup {
			stack = append(stack, token)
		} else if token.Type == CloseGroup {
			for len(stack) > 0 {
				t := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if t.Type == OpenGroup {
					break
				} else {
					queue = append(queue, t)
				}
			}
		} else {
			log.Fatal("invalid token")
		}
	}
	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return queue
}
