package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

type TwinResult[SO1 parser.ParseOutput, SO2 parser.ParseOutput] struct {
	One SO1
	Two SO2
}

// Twin initializes a parser that applies two sub-parsers sequentially.
func Twin[I comparable, SO1 parser.ParseOutput, SO2 parser.ParseOutput](
	s1 parser.Parser[I, SO1],
	s2 parser.Parser[I, SO2],
) parser.Parser[I, TwinResult[SO1, SO2]] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], TwinResult[SO1, SO2], parser.ParseError) {
		r := TwinResult[SO1, SO2]{}
		rest, so1, err := s1.Parse(input)
		if err != nil {
			return rest, r, err
		}
		r.One = so1

		rest, so2, err := s2.Parse(rest)
		if err != nil {
			return rest, r, err
		}

		r.Two = so2
		return rest, r, nil
	})
}
