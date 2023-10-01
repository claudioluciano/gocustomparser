package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestMany1SubFailureOnFirstApplication(t *testing.T) {
	subP := combinators.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	p := combinators.Many1(subP)

	i := strparse.NewCompleteInput("bbbbb")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
