package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Many0 initializes a parser that applies the given sub-parser several times.
func Many0[I comparable, SO parser.ParseOutput](sub parser.Parser[I, SO]) parser.Parser[I, []SO] {
	return many(sub, 0)
}

// Many1 initializes a parser that applies the given sub-parser several times.
// if the sub parser fails to parse and the count of application times is 0
// Many11 parser return an error.
func Many1[I comparable, SO parser.ParseOutput](sub parser.Parser[I, SO]) parser.Parser[I, []SO] {
	return many(sub, 1)
}

// many is the actual implementation of Many0/1.
func many[I comparable, SO parser.ParseOutput](sub parser.Parser[I, SO], min uint) parser.Parser[I, []SO] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], []SO, parser.ParseError) {
		count := 0
		output := make([]SO, 0)
		rest := input
		for {
			var o SO
			var err error

			rest, o, err = sub.Parse(rest)
			if err != nil {
				break
			}
			count++

			output = append(output, o)
		}

		if count < int(min) {
			return rest, output, &NotSatisfiedCountError{}
		}

		return rest, output, nil
	})
}

// NotSatisfiedCountError notifies the count of sub-parser success are not satisfied.
type NotSatisfiedCountError struct{}

// Error implements error interface.
func (e *NotSatisfiedCountError) Error() string {
	return "not satisfied the range of sub-parser succeeds"
}
