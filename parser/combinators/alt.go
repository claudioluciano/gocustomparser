package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Alt initializes a parser that applies all given parsers.
// if all of them are failed to parse, Alt() parser also returns an error.
// otherwise Alt() succeeds to parse.
func Alt[I comparable, O parser.ParseOutput](parsers ...parser.Parser[I, O]) parser.Parser[I, O] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O, parser.ParseError) {
		// subI holds the rest input in outer scope of for-statement.
		var subI parser.ParseInput[I]
		var subO O

		for _, subP := range parsers {
			var err parser.ParseError

			storedOffset, err := input.Seek(0, parser.SeekModeCurrent)
			if err != nil {
				return subI, subO, err
			}

			subI, subO, err = subP.Parse(input)
			if err == nil {
				return subI, subO, nil
			}

			// recover from subP's failure
			_, err = input.Seek(storedOffset, parser.SeekModeStart)
			if err != nil {
				return subI, subO, err
			}
		}

		return subI, subO, &AllParsersFailedError{}
	})
}

// AllParsersFailedError notifies all of given parsers are failed to parse.
type AllParsersFailedError struct{}

// Error implements error interface.
func (e *AllParsersFailedError) Error() string {
	return "all of given parser failed to parse"
}
