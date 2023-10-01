package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestSatisfyFailure(t *testing.T) {
	p := combinators.Satisfy(func(ch rune) bool { return ch == 'a' })

	i := strparse.NewCompleteInput("bbc")
	_, _, err := p.Parse(i)

	assert.Error(t, err)
}
