package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// ManyMinMax initializes a parser that applies the given sub-parser several times.
// It fails if the sub parser does not succeed at least min times.
// It also fails if the sub parser does succeed over max times.
func ManyMinMax[I comparable, SO parser.ParseOutput](sub parser.Parser[I, SO], min uint, max uint) parser.Parser[I, []SO] {
	return manyMinMax(sub, min, max)
}

// manyMinMax is the actual implementation of ManyMinMax.
func manyMinMax[I comparable, SO parser.ParseOutput](sub parser.Parser[I, SO], min uint, max uint) parser.Parser[I, []SO] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], []SO, parser.ParseError) {
		count := 0
		output := make([]SO, 0)
		for {
			var o SO
			var err error

			if count >= int(max) {
				return input, output, &NotSatisfiedCountError{}
			}

			input, o, err = sub.Parse(input)
			if err != nil {
				break
			}
			count++

			output = append(output, o)
		}

		if count < int(min) {
			return input, output, &NotSatisfiedCountError{}
		}

		return input, output, nil
	})
}
