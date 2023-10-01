package byteparse_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestTagFailureWithShorterInput(t *testing.T) {
	p := byteparse.Tag([]byte("Golang"))

	b := byteparse.NewCompleteInput([]byte("Go"))
	_, _, err := p.Parse(b)
	assert.Error(t, err)
}

func TestTagFailure(t *testing.T) {
	p := byteparse.Tag([]byte("Clang"))

	b := byteparse.NewCompleteInput([]byte("Dlang"))
	_, _, err := p.Parse(b)
	assert.Error(t, err)
}
