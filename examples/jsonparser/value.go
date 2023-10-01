package main

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonValue interface{}

// parseJSONValue parses the json value.
// value := whitespace_opt (string | number | array) whitespace_opt
func parseJSONValue() parser.Parser[byte, jsonValue] {
	begin := parseJSONWhitespace()
	contents := combinators.Lazy(func() parser.Parser[byte, jsonValue] {
		return combinators.Alt(
			parseJSONStringValue(),
			parseJSONNumberValue(),
			parseJSONBooleanValue(),
			parseJSONArrayValue(),
			parseJSONObjectValue(),
		)
	})
	end := parseJSONWhitespace()
	p := combinators.Delimited(begin, contents, end)
	return p
}
