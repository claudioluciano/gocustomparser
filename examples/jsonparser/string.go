package main

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonStringValue string

func parseJSONStringValue() parser.Parser[byte, jsonValue] {
	quote := combinators.Satisfy(func(ch byte) bool { return ch == '"' })
	contents := combinators.Many0(combinators.Satisfy(func(ch byte) bool { return ch != '"' }))

	p := combinators.Map(combinators.Delimited(quote, contents, quote), func(s []byte) (jsonValue, error) {
		return jsonStringValue(s), nil
	})

	return p
}
