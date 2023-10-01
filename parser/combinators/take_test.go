package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestTakeFailure(t *testing.T) {
	subP := combinators.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	p := combinators.Take(5, subP)

	i := strparse.NewCompleteInput("aaaab")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
