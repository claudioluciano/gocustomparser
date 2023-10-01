package strparse_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestTagFailureWithShorterInput(t *testing.T) {
	p := strparse.Tag("Golang")

	i := strparse.NewCompleteInput("Go")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestTagFailure(t *testing.T) {
	p := strparse.Tag("Clang")

	i := strparse.NewCompleteInput("Dlang")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
