package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Take initializes a parser that applies sub-parser count times.
func Take[I comparable, SO parser.ParseOutput](count uint, sub parser.Parser[I, SO]) parser.Parser[I, []SO] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], []SO, parser.ParseError) {
		output := make([]SO, count)

		var o SO
		var err error
		for i := uint(0); i < count; i++ {
			input, o, err = sub.Parse(input)
			if err != nil {
				return input, output, err
			}

			output[i] = o
		}

		return input, output, nil
	})
}
