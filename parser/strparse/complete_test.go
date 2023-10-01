package strparse

import (
	"testing"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	i := NewCompleteInput("12")

	buf := make([]rune, 1)
	n, err := i.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, rune('1'), buf[0])

	n, err = i.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, rune('2'), buf[0])

	n, err = i.Read(buf)
	assert.Error(t, err)
	assert.Equal(t, 0, n)
}

func TestSeek(t *testing.T) {
	i := NewCompleteInput("12")

	buf := make([]rune, 1)
	n, err := i.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, rune('1'), buf[0])

	n, err = i.Seek(0, parser.SeekModeStart)
	assert.NoError(t, err)
	assert.Equal(t, 0, n)

	n, err = i.Read(buf)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, rune('1'), buf[0])
}
