package combinators

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Branches initializes a parser that receives multiple syntax-rules and determine one of them.
// In almost cases, user can enumerate all syntax rules before starting parsing.
// so branches receives map. We recommend you to initialize the map at once and refer multiple times.
// It may be efficient.
// if no applicable rule exists in the rules, Branches() parser returns an error.
// if all of them are failed to parse, Branches() parser also returns an error.
func Branches[I comparable, O parser.ParseOutput](rules map[I]parser.Parser[I, O]) parser.Parser[I, O] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O, parser.ParseError) {
		var o O

		storedOffset, err := input.Seek(0, parser.SeekModeCurrent)
		if err != nil {
			return input, o, err
		}

		buf := make([]I, 1)
		n, err := input.Read(buf)
		if err != nil || n < 1 {
			return input, o, err
		}

		subP, ok := rules[buf[0]]
		if !ok {
			return input, o, &parser.NoLeftInputToParseError{}
		}

		// recover the consumed head of the input stream
		_, err = input.Seek(storedOffset, parser.SeekModeStart)
		if err != nil {
			return input, o, err
		}

		return subP.Parse(input)
	})
}

// ApplicableRuleIsNotFoundError notifies all of given parsers don't match the head of the input.
type ApplicableRuleIsNotFoundError[I comparable] struct {
	// actual is the given element
	actual I
}

// Error implements error interface.
func (e *ApplicableRuleIsNotFoundError[E]) Error() string {
	return fmt.Sprintf("all of given parser cannot start parsing on '%v'", e.actual)
}
