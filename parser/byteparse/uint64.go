package byteparse

import (
	"encoding/binary"

	"github.com/claudioluciano/gocustomparser/parser"
)

// Uint64 initializes a parser that parse 64-bit unsigned integer.
// user can determine the behavior of this parser by giving byteorder what you want to use.
func UInt64(byteorder binary.ByteOrder) parser.Parser[byte, uint64] {
	return parser.NewParser(func(input parser.ParseInput[byte]) (parser.ParseInput[byte], uint64, parser.ParseError) {
		buf := make([]byte, 8)

		n, err := input.Read(buf)
		if err != nil || n < 8 {
			return input, 0, &parser.NoLeftInputToParseError{}
		}

		v := byteorder.Uint64(buf)
		return input, v, nil
	})
}
