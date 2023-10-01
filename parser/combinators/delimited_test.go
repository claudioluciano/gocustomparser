package combinators_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestDelimitedBeginFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinators.Delimited(begin, contents, end)
	i := strparse.NewCompleteInput("'12345\"")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestDelimitedContentsFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinators.Delimited(begin, contents, end)
	i := strparse.NewCompleteInput("\"abcde\"")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestDelimitedEndFailure(t *testing.T) {
	begin := strparse.Rune('"')
	contents := strparse.Digit1()
	end := strparse.Rune('"')

	p := combinators.Delimited(begin, contents, end)
	i := strparse.NewCompleteInput("\"12345'")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
