package main

type Type int64

const (
	Literal Type = iota
	Wildcard
	Concat
	Union
	Optional
	Plus
	Star
	Group
	CharSet
	Repeat
	Skip
)

var Special = map[byte]bool{
	'*': true,
	'+': true,
	'?': true,
	'|': true,
	'(': true,
	')': true,
	'{': true,
	'}': true,
}

var Operation = map[byte]bool{
	'*': true,
	'+': true,
	'?': true,
	'|': true,
}

var CharType = map[byte]Type{
	'.':  Wildcard,
	'*':  Star,
	'+':  Plus,
	'?':  Optional,
	'|':  Union,
	'(':  Group,
	')':  Group,
	'[':  CharSet,
	']':  CharSet,
	'{':  Repeat,
	'}':  Repeat,
	'\\': Skip,
}

func (t Type) Precedence() int {
	switch t {
	case Concat:
		return 1
	case Union:
		return 2
	case Optional:
		return 3
	case Plus:
		return 3
	case Star:
		return 3
	}
	return -1
}

type Token struct {
	Value     string
	Operation Type
}
