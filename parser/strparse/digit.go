package strparse

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

// Digit1 initializes a parser that follows (0-9)+ syntax rule.
func Digit1() parser.Parser[rune, string] {
	return parser.NewParser(func(input parser.ParseInput[rune]) (parser.ParseInput[rune], string, parser.ParseError) {
		p := combinators.Many1(combinators.Satisfy(isDigit))
		i, o, err := p.Parse(input)
		if err != nil {
			return i, "", err
		}

		return i, string(o), nil
	})
}

// isDigit checks the given rune is in the range of unicode digits.
func isDigit(ch rune) bool {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}
