package main

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONValue(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte("     12345     "))
	p := parseJSONValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, jsonNumberValue(12345), o)
}
