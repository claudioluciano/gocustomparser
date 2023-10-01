package byteparse

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// Uint8 initializes a parser that parse 8-bit unsigned integer.
func UInt8() parser.Parser[byte, uint8] {
	return parser.NewParser(func(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint8, parser.ParseError) {
		buf := make([]byte, 1)

		n, err := input.Read(buf)
		if err != nil || n < 1 {
			return input, 0, &parser.NoLeftInputToParseError{}
		}

		v := buf[0]
		return input, v, nil
	})
}
