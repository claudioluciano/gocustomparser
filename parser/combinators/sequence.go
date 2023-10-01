package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Sequence initializes a parser that applies a sequence of sub-parsers.
func Sequence[I comparable, SO parser.ParseOutput](subs []parser.Parser[I, SO]) parser.Parser[I, []SO] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], []SO, parser.ParseError) {
		result := make([]SO, len(subs))
		rest := input

		for i := range subs {
			var o SO
			var err error
			rest, o, err = subs[i].Parse(rest)
			if err != nil {
				return rest, result, err
			}

			result[i] = o
		}

		return rest, result, nil
	})
}
