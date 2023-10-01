package main

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonBooleanValue bool

// parseJSONBooleanValue parses the boolean literal.
// boolean := true | false
func parseJSONBooleanValue() parser.Parser[byte, jsonValue] {
	const trueSig = "true"
	const falseSig = "false"
	trueP := byteparse.Tag([]byte(trueSig))
	falseP := byteparse.Tag([]byte(falseSig))
	p := combinators.Map(combinators.Alt(trueP, falseP), func(s []byte) (jsonValue, error) {
		switch string(s) {
		case trueSig:
			return jsonBooleanValue(true), nil
		case falseSig:
			return jsonBooleanValue(false), nil
		default:
			// maybe unreachable
			return nil, fmt.Errorf("unexpected bytes tag '%s'", s)
		}
	})

	return p
}
