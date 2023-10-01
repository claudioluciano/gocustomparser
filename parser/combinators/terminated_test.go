package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestTerminatedPredecessorFailure(t *testing.T) {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinators.Terminated(predecessor, successor)

	i := strparse.NewCompleteInput("/a")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestTerminatedSuccessorFailure(t *testing.T) {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinators.Terminated(predecessor, successor)

	i := strparse.NewCompleteInput("*b")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
