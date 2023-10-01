package main

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestParseJSONBooleanValue_True(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte("true"))
	p := parseJSONBooleanValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, jsonBooleanValue(true), o)
}

func TestParseJSONBooleanValue_False(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte("false"))
	p := parseJSONBooleanValue()
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, jsonBooleanValue(false), o)
}
