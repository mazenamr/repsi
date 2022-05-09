package parser

type TokenType int

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

var CharTokenType = map[byte]TokenType{
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

var IsOperator = map[TokenType]bool{
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
