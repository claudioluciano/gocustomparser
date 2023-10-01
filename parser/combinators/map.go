package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Map initializes a parser that applies a sub-parser and give the it's output to fn.
func Map[I comparable, SO parser.ParseOutput, O parser.ParseOutput](sub parser.Parser[I, SO], fn func(SO) (O, error)) parser.Parser[I, O] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O, parser.ParseError) {
		var o O

		i, so, err := sub.Parse(input)
		if err != nil {
			return i, o, err
		}

		o, err = fn(so)
		return i, o, err
	})
}
