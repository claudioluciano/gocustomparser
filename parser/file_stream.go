package parser

import (
	"io"
)

// IOReadSeeker is the stream input.
// We should require io.Seeker because some parsers try to recover from failures.
type IOReadSeeker struct {
	r io.ReadSeeker
}

// NewIOReadSeeker initializes an IOReadSeeker.
func NewIOReadSeeker(r io.ReadSeeker) *IOReadSeeker {
	return &IOReadSeeker{r}
}

// Read implements ParseInput interface.
func (r *IOReadSeeker) Read(buf []byte) (int, error) {
	return r.r.Read(buf)
}

// Seek implements ParseInput interface.
func (r *IOReadSeeker) Seek(n int, mode SeekMode) (int, error) {
	v, err := r.r.Seek(int64(n), int(mode))
	return int(v), err
}
