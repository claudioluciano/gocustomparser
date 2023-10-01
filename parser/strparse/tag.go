package strparse

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Tag initializes a parser that checks the input starts with the tag prefix.
func Tag(tag string) parser.Parser[rune, string] {
	return parser.NewParser(func(input parser.ParseInput[rune]) (parser.ParseInput[rune], string, parser.ParseError) {
		t := []rune(tag)
		buf := make([]rune, len(tag))

		n, err := input.Read(buf)
		if err != nil || n < len(tag) {
			return input, "", &parser.NoLeftInputToParseError{}
		}

		unmatched := !hasPrefix(buf, t)
		if unmatched {
			return input, "", &UnexpectedPrefixError{expected: tag}
		}

		return input, tag, nil
	})
}

// UnexpectedPrefixError notifies the prefix of the given input is unexpected.
type UnexpectedPrefixError struct {
	expected string
}

// Error implements error interface.
func (e *UnexpectedPrefixError) Error() string {
	return fmt.Sprintf("expected \"%s\" prefix", e.expected)
}
