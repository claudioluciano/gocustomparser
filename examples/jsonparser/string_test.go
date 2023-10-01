package main

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONStringValue(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte("\"Hello World\""))
	p := parseJSONStringValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, jsonStringValue("Hello World"), o)
}
