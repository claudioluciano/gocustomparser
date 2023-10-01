package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestManyMinMaxFailure(t *testing.T) {
	subP := combinators.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})

	i := strparse.NewCompleteInput("aabbb")
	p := combinators.ManyMinMax(subP, 3, 5)
	_, _, err := p.Parse(i)
	assert.Error(t, err)

	i = strparse.NewCompleteInput("aaaaaabbb")
	_, _, err = p.Parse(i)
	assert.Error(t, err)
}
