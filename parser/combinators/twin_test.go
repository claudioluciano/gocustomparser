package combinators_test

import (
	"fmt"
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestTwinSubParserFailure(t *testing.T) {
	a := strparse.Rune('a')
	b := combinators.Map(strparse.Rune('b'), func(o rune) (string, error) {
		return "", fmt.Errorf("error")
	})

	p := combinators.Twin(a, b)

	i := strparse.NewCompleteInput("ab")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
