package combinators

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Satisfy initializes a parser that checks the head of the input satisfies the predicate.
func Satisfy[I comparable](pred Predicate[I]) parser.Parser[I, I] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], I, parser.ParseError) {
		var e I

		storedOffset, err := input.Seek(0, parser.SeekModeCurrent)
		if err != nil {
			return input, e, err
		}

		buf := make([]I, 1)
		n, err := input.Read(buf)
		if err != nil || n < 1 {
			return input, e, &parser.NoLeftInputToParseError{}
		}

		notSatisfied := !pred(buf[0])
		if notSatisfied {
			// recover the consumed head of the input stream.
			input.Seek(storedOffset, parser.SeekModeStart)
			return input, e, &NotSatisfiedError[I]{actual: e}
		}

		return input, buf[0], nil
	})
}

// Predicate is the condition that satisfyParser uses for consuming one element.
type Predicate[I comparable] func(input I) bool

// NotsatisfiedError notifies that the given predicate is not satisfied.
type NotSatisfiedError[I comparable] struct {
	// actual is the given element that satisfyParser consumed
	actual I
}

// Error implements error interface
func (e *NotSatisfiedError[E]) Error() string {
	return fmt.Sprintf("predicate was not satisfied on '%+v'", e.actual)
}
