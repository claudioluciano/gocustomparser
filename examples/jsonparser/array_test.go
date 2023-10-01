package main

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONArrayValue(t *testing.T) {
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonStringValue("foo"),
			jsonStringValue("bar"),
			jsonStringValue("baz"),
		},
		length: 3,
	}

	i := byteparse.NewCompleteInput([]byte(`["foo", "bar", "baz"]`))
	p := parseJSONArrayValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONArrayValueWith2d(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte(`[["a", "b"], ["c", "d"], ["e", "f"]]`))
	expected := jsonArrayValue{
		elements: []jsonValue{
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("a"),
					jsonStringValue("b"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("c"),
					jsonStringValue("d"),
				},
				length: 2,
			},
			jsonArrayValue{
				elements: []jsonValue{
					jsonStringValue("e"),
					jsonStringValue("f"),
				},
				length: 2,
			},
		},
		length: 3,
	}
	p := parseJSONArrayValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}
