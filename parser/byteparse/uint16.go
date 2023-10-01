package byteparse

import (
	"encoding/binary"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Uint16 initializes a parser that parse 16-bit unsigned integer.
// user can determine the behavior of this parser by giving byteorder what you want to use.
func UInt16(byteorder binary.ByteOrder) parser.Parser[byte, uint16] {
	return parser.NewParser(func(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint16, parser.ParseError) {
		buf := make([]byte, 2)

		n, err := input.Read(buf)
		if err != nil || n < 2 {
			return input, 0, &parser.NoLeftInputToParseError{}
		}

		v := byteorder.Uint16(buf)
		return input, v, nil
	})
}
