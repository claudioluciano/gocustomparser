package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/stretchr/testify/assert"
)

func TestBranchesFailure(t *testing.T) {
	m := make(map[byte]parser.Parser[byte, uint])
	m[0x00] = combinators.Map(byteparse.UInt8(), func(v uint8) (uint, error) { return uint(v), nil })
	m[0x01] = combinators.Map(byteparse.UInt8(), func(v uint8) (uint, error) { return uint(v), nil })
	p := combinators.Branches(m)

	i := byteparse.NewCompleteInput([]byte{0x02})
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
