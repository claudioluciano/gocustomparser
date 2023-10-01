package combinators_test

import (
	"fmt"
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestMapSubParserFailure(t *testing.T) {
	subP := strparse.Rune('a')
	p := combinators.Map(subP, func(ch rune) (bool, error) { return ch == 'a', nil })

	i := strparse.NewCompleteInput("bc")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestMapFnFailure(t *testing.T) {
	subP := strparse.Rune('a')
	p := combinators.Map(subP, func(ch rune) (bool, error) { return false, fmt.Errorf("") })

	i := strparse.NewCompleteInput("a")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
