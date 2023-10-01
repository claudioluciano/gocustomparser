package main

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONObjectValue_Empty(t *testing.T) {
	expected := jsonObjectValue{
		Fields: make([]jsonObjectField, 0),
	}
	i := byteparse.NewCompleteInput([]byte(`{ }`))
	p := parseJSONObjectValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectValue(t *testing.T) {
	expected := jsonObjectValue{
		Fields: []jsonObjectField{
			{
				Name:  "A",
				Value: jsonStringValue("B"),
			},
			{
				Name:  "C",
				Value: jsonStringValue("D"),
			},
		},
	}
	i := byteparse.NewCompleteInput([]byte(`{"A":"B","C":"D"}`))
	p := parseJSONObjectValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectField(t *testing.T) {
	expected := jsonObjectField{
		Name:  "foo",
		Value: jsonStringValue("bar"),
	}
	i := byteparse.NewCompleteInput([]byte(`"foo" : "bar"`))
	p := parseJSONObjectField()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}

func TestParseJSONObjectFieldList(t *testing.T) {
	expected := []jsonObjectField{
		{
			Name:  "n1",
			Value: jsonStringValue("v1"),
		},
		{
			Name:  "n2",
			Value: jsonStringValue("v2"),
		},
	}
	i := byteparse.NewCompleteInput([]byte(`"n1" : "v1", "n2" : "v2"`))
	p := parseJSONObjectFieldList()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, expected, o)
}
