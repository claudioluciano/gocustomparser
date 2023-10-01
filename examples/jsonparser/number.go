package main

import (
	"strconv"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonNumberValue int

func parseJSONNumberValue() parser.Parser[byte, jsonValue] {
	digit1 := combinators.Many1(combinators.Satisfy(func(b byte) bool {
		return '0' <= b && b <= '9'
	}))
	p := combinators.Map(digit1, func(s []byte) (jsonValue, error) {
		v, err := strconv.ParseInt(string(s), 10, 64)
		return jsonNumberValue(v), err
	})

	return p
}
