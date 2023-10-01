package byteparse

import (
	"encoding/binary"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Uint32 initializes a parser that parse 32-bit unsigned integer.
// user can determine the behavior of this parser by giving byteorder what you want to use.
func UInt32(byteorder binary.ByteOrder) parser.Parser[byte, uint32] {
	return parser.NewParser(func(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint32, parser.ParseError) {
		buf := make([]byte, 4)

		n, err := input.Read(buf)
		if err != nil || n < 4 {
			return input, 0, &parser.NoLeftInputToParseError{}
		}

		v := byteorder.Uint32(buf)
		return input, v, nil
	})
}
