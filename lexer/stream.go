package lexer

import "github.com/claudioluciano/gocustomparser/parser"

type TokenStream struct {
	tokens []*Token
	offset int
}

func NewTokenStream(tokens []*Token) *TokenStream {
	return &TokenStream{
		tokens: tokens,
		offset: 0,
	}
}

func (c *TokenStream) Read(buf []*Token) (int, error) {
	if c.offset >= len(c.tokens) {
		return 0, &parser.NoLeftInputToParseError{}
	}

	copy(buf, c.tokens[c.offset:])
	c.offset += len(buf)

	return len(buf), nil
}

func (c *TokenStream) Seek(n int, mode parser.SeekMode) (int, error) {
	switch mode {
	case parser.SeekModeStart:
		if n >= len(c.tokens) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset = n
		return c.offset, nil
	case parser.SeekModeCurrent:
		if c.offset+n >= len(c.tokens) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset += n
		return c.offset, nil
	default:
		panic("given seek mode is not supported")
	}
}
