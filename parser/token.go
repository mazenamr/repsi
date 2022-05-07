package parser

type TokenType int64

const (
	Literal TokenType = iota
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

var Special = map[byte]bool{
	'*': true,
	'+': true,
	'?': true,
	'|': true,
	'(': true,
	')': true,
	'{': true,
}

var CharType = map[byte]TokenType{
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

var Operator = map[TokenType]bool{
	Union:    true,
	Concat:   true,
	Repeat:   true,
	Optional: true,
	Plus:     true,
	Star:     true,
}

var Precedence = map[TokenType]int{
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
	Type  TokenType
}
