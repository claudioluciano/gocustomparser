package byteparse

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// CompleteInput holds the whole bytes.
type CompleteInput struct {
	bytes  []byte
	offset int
}

// NewCompleteInput initialiizes a CompleteInput.
func NewCompleteInput(bytes []byte) *CompleteInput {
	return &CompleteInput{bytes: bytes}
}

// Read implements parser.ParseInput interface.
func (c *CompleteInput) Read(buf []byte) (int, error) {
	if c.offset >= len(c.bytes) {
		return 0, &parser.NoLeftInputToParseError{}
	}

	copy(buf, c.bytes[c.offset:])
	c.offset += len(buf)

	return len(buf), nil
}

// Seek implements parser.ParseInput interface.
func (c *CompleteInput) Seek(n int, mode parser.SeekMode) (int, error) {
	switch mode {
	case parser.SeekModeStart:
		if n >= len(c.bytes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset = n
		return c.offset, nil
	case parser.SeekModeCurrent:
		if c.offset+n >= len(c.bytes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset += n
		return c.offset, nil
	default:
		panic("given seek mode is not supported")
	}
}
