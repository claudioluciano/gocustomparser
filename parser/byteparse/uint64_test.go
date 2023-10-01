package byteparse_test

import (
	"encoding/binary"
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/stretchr/testify/assert"
)

func TestUInt64Failure(t *testing.T) {
	i := byteparse.NewCompleteInput([]byte{})
	p := byteparse.UInt64(binary.BigEndian)
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
