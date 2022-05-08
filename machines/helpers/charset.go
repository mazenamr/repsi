package helpers

import (
	"repsi/consts"
	"sort"
)

func Match(t string, s string) bool {
	if s == consts.WildcardToken {
		return true
	}
	if len(s) > 1 {
		if t[0] == '[' && t[len(t)-1] == ']' {
			charset := ""
			if t[1] != '^' {
				charset = t[1 : len(t)-1]
			} else {
				charset = t[2 : len(t)-1]
			}
			chars := ExpandCharset(charset)
			if len(chars) == 0 {
				return false
			}
			for _, c := range chars {
				if c == t {
					return t[1] != '^'
				}
			}
		}
	}
	return t == s
}

func ExpandCharset(s string) []string {
	chars := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			if i+1 < len(s) {
				for c := s[i-1]; c <= s[i+1]; c++ {
					chars = append(chars, string(c))
				}
				i++
			}
		} else {
			chars = append(chars, string(s[i]))
		}
	}
	return chars
}

func GenerateCharacterSet(s []string) string {
	keys := make(map[string]bool)
	for _, k := range s {
		if k == consts.WildcardToken {
			return consts.WildcardToken
		}
		if len(k) == 1 {
			keys[k] = true
		}
	}

	var chars []string
	for k := range keys {
		chars = append(chars, k)
	}
	sort.Strings(chars)

	c := ""
	for _, char := range chars {
		c += char
	}

	if len(c) == 1 {
		return string(c)
	}

	charset := "["
	for i := 0; i < len(c); i++ {
		charset += string(c[i])
		conc := false
		for i+1 < len(c) && c[i+1] == c[i]+1 {
			conc = true
			i++
		}
		if conc {
			charset += "-"
			charset += string(c[i])
		}
	}
	charset += "]"
	return charset
}
