package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestAltChildrenFailure(t *testing.T) {
	p1 := strparse.Rune('a')
	p2 := strparse.Rune('b')
	p := combinators.Alt(p1, p2)

	i := strparse.NewCompleteInput("c")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
