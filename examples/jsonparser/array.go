package main

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonArrayValue struct {
	elements []jsonValue
	length   int
}

func parseJSONArrayValue() parser.Parser[byte, jsonValue] {
	begin := combinators.Satisfy(func(b byte) bool {
		return b == '['
	})
	end := combinators.Satisfy(func(b byte) bool {
		return b == ']'
	})
	separator := combinators.Satisfy(func(b byte) bool {
		return b == ','
	})
	element := parseJSONValue()
	contents := combinators.Separated1(element, separator)
	p := combinators.Map(combinators.Delimited(begin, contents, end), func(v []jsonValue) (jsonValue, error) {
		return jsonArrayValue{elements: v, length: len(v)}, nil
	})
	return p
}
