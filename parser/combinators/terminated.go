package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Terminated initializes a parser that applies given parsers but discards successor's output.
func Terminated[
	I comparable,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
](predecessor parser.Parser[I, O1], successor parser.Parser[I, O2],
) parser.Parser[I, O1] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O1, parser.ParseError) {
		rest, o1, err := predecessor.Parse(input)
		if err != nil {
			return rest, o1, err
		}

		rest, _, err = successor.Parse(rest)
		return rest, o1, err
	})
}
