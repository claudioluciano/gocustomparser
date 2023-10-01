package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestSequenceSubParserFailure(t *testing.T) {
	a := strparse.Rune('a')
	b := strparse.Rune('b')
	c := strparse.Rune('c')
	p := combinators.Sequence([]parser.Parser[rune, rune]{a, b, c})

	i := strparse.NewCompleteInput("abd")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
