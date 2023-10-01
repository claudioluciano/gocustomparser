package byteparse

import (
	"bytes"
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Tag initializes a parser that checks the input starts with the tag prefix.
func Tag(tag []byte) parser.Parser[byte, []byte] {
	return parser.NewParser(func(input parser.ParseInput[byte]) (parser.ParseInput[byte], []byte, parser.ParseError) {
		storedOffset, err := input.Seek(0, parser.SeekModeCurrent)
		if err != nil {
			return input, nil, err
		}

		buf := make([]byte, len(tag))

		n, err := input.Read(buf)
		if err != nil || n < len(tag) {
			// recover the consumed head of the input stream.
			input.Seek(storedOffset, parser.SeekModeStart)
			return input, nil, &parser.NoLeftInputToParseError{}
		}

		unmatched := !bytes.HasPrefix(buf, tag)
		if unmatched {
			// recover the consumed head of the input stream.
			input.Seek(storedOffset, parser.SeekModeStart)
			return input, nil, &UnexpectedPrefixError{expected: tag}
		}
		return input, tag, nil
	})
}

// UnexpectedPrefixError notifies the prefix of the given input is unexpected.
type UnexpectedPrefixError struct {
	expected []byte
}

// Error implements error interface.
func (e *UnexpectedPrefixError) Error() string {
	return fmt.Sprintf("expected \"%s\" prefix", e.expected)
}
