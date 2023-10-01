package strparse

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

// CompleteInput holds the whole runes.
type CompleteInput struct {
	runes  []rune
	offset int
}

// NewCompleteInput initialiizes a CompleteInput.
func NewCompleteInput(s string) *CompleteInput {
	return &CompleteInput{runes: []rune(s)}
}

// Read implements parser.ParseInput interface.
func (c *CompleteInput) Read(buf []rune) (int, error) {
	if c.offset >= len(c.runes) {
		return 0, &parser.NoLeftInputToParseError{}
	}

	copy(buf, c.runes[c.offset:])
	c.offset += len(buf)

	return len(buf), nil
}

// Seek implements parser.ParseInput interface.
func (c *CompleteInput) Seek(n int, mode parser.SeekMode) (int, error) {
	switch mode {
	case parser.SeekModeStart:
		if n >= len(c.runes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset = n
		return c.offset, nil
	case parser.SeekModeCurrent:
		if c.offset+n >= len(c.runes) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset += n
		return c.offset, nil
	default:
		panic("given seek mode is not supported")
	}
}
