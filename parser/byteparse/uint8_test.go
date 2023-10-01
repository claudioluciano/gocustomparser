package byteparse_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestUInt8Failure(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte{})
	p := byteparse.UInt8()
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
