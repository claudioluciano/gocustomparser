package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Preceded initializes a parser that applies given parsers but discards predecessor's output.
func Preceded[
	I comparable,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
](predecessor parser.Parser[I, O1], successor parser.Parser[I, O2],
) parser.Parser[I, O2] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O2, parser.ParseError) {
		var o2 O2
		rest, _, err := predecessor.Parse(input)
		if err != nil {
			return rest, o2, err
		}

		return successor.Parse(rest)
	})
}
