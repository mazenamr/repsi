package parser

type tokenType int64

const (
	Literal tokenType = iota
	Wildcard
	CharSet
	Union
	Concat
	Repeat
	Optional
	Plus
	Star
	OpenGroup
	CloseGroup
)

var charTokenType = map[byte]tokenType{
	'.': Wildcard,
	'[': CharSet,
	'|': Union,
	'{': Repeat,
	'?': Optional,
	'+': Plus,
	'*': Star,
	'(': OpenGroup,
	')': CloseGroup,
}

var isOperator = map[tokenType]bool{
	Union:    true,
	Concat:   true,
	Repeat:   true,
	Optional: true,
	Plus:     true,
	Star:     true,
}

var precedence = map[tokenType]int{
	OpenGroup: 0,
	Union:     1,
	Concat:    2,
	Repeat:    3,
	Optional:  3,
	Plus:      3,
	Star:      3,
}

type Token struct {
	Value string
	Type  tokenType
}
