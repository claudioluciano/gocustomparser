package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestSeparated1FirstElementFailure(t *testing.T) {
	element := strparse.Digit1()
	separator := strparse.Rune('|')
	p := combinators.Separated1(element, separator)

	i := strparse.NewCompleteInput("abc|123|123")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestSeparated1WithOneElement(t *testing.T) {
	element := strparse.Digit1()
	separator := strparse.Rune('|')
	p := combinators.Separated1(element, separator)

	i := strparse.NewCompleteInput("123")
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(o))
	assert.Equal(t, "123", o[0])
}
