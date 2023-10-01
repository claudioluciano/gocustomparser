package main

import (
	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type jsonObjectValue struct {
	Fields []jsonObjectField
}

// parseJSONObjectValue parses a JSON object.
func parseJSONObjectValue() parser.Parser[byte, jsonValue] {
	begin := byteparse.Tag([]byte("{"))

	end := combinators.Preceded(parseJSONWhitespace(), byteparse.Tag([]byte("}")))

	emptyObject := combinators.Map(combinators.Twin(begin, end), func(res combinators.TwinResult[[]byte, []byte]) (jsonValue, error) {
		return jsonObjectValue{Fields: make([]jsonObjectField, 0)}, nil
	})

	fields := parseJSONObjectFieldList()
	rawObject := combinators.Delimited(begin, fields, end)
	object := combinators.Map(rawObject, func(fields []jsonObjectField) (jsonValue, error) {
		return jsonObjectValue{Fields: fields}, nil
	})

	p := combinators.Alt(emptyObject, object)
	return p
}

type jsonObjectField struct {
	Name  string
	Value jsonValue
}

func parseJSONObjectFieldList() parser.Parser[byte, []jsonObjectField] {
	ws := parseJSONWhitespace()
	comma := combinators.Delimited(ws, byteparse.Tag([]byte(",")), ws)
	p := combinators.Separated1(parseJSONObjectField(), comma)
	return p
}

// parseJSONObjectField parses a field of the JSON object
// object_field := string whitespace ":" value
func parseJSONObjectField() parser.Parser[byte, jsonObjectField] {
	rawName := parseJSONStringValue()
	ws := parseJSONWhitespace()
	nameWS := combinators.Delimited(ws, rawName, ws)
	colon := byteparse.Tag([]byte(":"))
	nameP := combinators.Terminated(nameWS, colon)
	p := combinators.Twin(nameP, parseJSONValue())

	return combinators.Map(p, func(res combinators.TwinResult[jsonValue, jsonValue]) (jsonObjectField, error) {
		name, _ := res.One.(jsonStringValue)
		return jsonObjectField{
			Name:  string(name),
			Value: res.Two,
		}, nil
	})
}
