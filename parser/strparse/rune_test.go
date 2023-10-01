package strparse_test

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser/strparse"
	"github.com/stretchr/testify/assert"
)

func TestRuneFailure(t *testing.T) {
	p := strparse.Rune('a')

	i := strparse.NewCompleteInput("bbc")
	_, _, err := p.Parse(i)
	assert.Error(t, err)
}

func TestRuneOnMuitiBytes(t *testing.T) {
	p := strparse.Rune('あ')

	i := strparse.NewCompleteInput("あいう")
	_, o, err := p.Parse(i)
	assert.NoError(t, err)
	assert.Equal(t, 'あ', o)
}
