package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestPrecededPredecessorFailure(t *testing.T) {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinators.Preceded(predecessor, successor)

	i := strparse.NewCompleteInput("/a")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestPrecededSuccessorFailure(t *testing.T) {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinators.Preceded(predecessor, successor)

	i := strparse.NewCompleteInput("*b")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
