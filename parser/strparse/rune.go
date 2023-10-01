package strparse

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

// Rune initializes a parser that consumes one rune.
// It's just a specialized parser from combinators.Satisfy().
func Rune(expected rune) parser.Parser[rune, rune] {
	return parser.NewParser(func(input parser.ParseInput[rune]) (parser.ParseInput[rune], rune, parser.ParseError) {
		p := combinators.Satisfy(func(ch rune) bool {
			return ch == expected
		})
		return p.Parse(input)
	})
}

// UnexpectedRuneError notifies the head of the given input is unexpected.
type UnexpectedRuneError struct {
	actual   rune
	expected rune
}

// Error implements error interface.
func (e *UnexpectedRuneError) Error() string {
	return fmt.Sprintf("expected '%c' but got '%c'", e.expected, e.actual)
}
