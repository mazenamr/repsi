package parser

import (
	"repsi/consts"
	"repsi/nfa"
)

func Parse(s string) *nfa.Machine {
	tokens := Preprocess(s)
	if !CheckBrakets(tokens) {
		panic("invalid regex")
	}
	queue := Postfix(tokens)
	stack := make([]*nfa.Machine, 0, len(queue))
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
				panic("invalid regex")
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
				panic("invalid regex")
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
			} else if i < len(s) && s[i] == '}' {
				tokens = append(tokens, &Token{Value: string(min), Type: Repeat})
			} else {
				panic("invalid regex")
			}
		case '\\':
			i++
			tokens = append(tokens, &Token{Value: string(s[i]), Type: Literal})
		default:
			tokens = append(tokens, &Token{Value: string(s[i]), Type: Literal})
		}

		if i+1 < len(s) && !Operator[CharType[s[i+1]]] {
			tokens = append(tokens, &Token{Type: Concat})
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
					continue
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
			panic("invalid token")
		}
	}
	for len(stack) > 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return queue
}
