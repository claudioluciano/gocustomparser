package strparse_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestDigit1Failure(t *testing.T) {
	p := strparse.Digit1()

	i := strparse.NewCompleteInput("aabbccdd11223344")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}
