package main

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

func parseJSONWhitespace() parser.Parser[byte, []byte] {
	return combinators.Many0(combinators.Satisfy(func(ch byte) bool {
		return ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t'
	}))
}
