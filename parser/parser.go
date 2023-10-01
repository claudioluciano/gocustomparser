package parser

import "io"

type Parser[I comparable, O ParseOutput] struct {
	fn  ParserFn[I, O]
	err ParseError
}

// ParserFn is an abstract parser.
type ParserFn[I comparable, O ParseOutput] func(input ParseInput[I]) (ParseInput[I], O, ParseError)

// ParseInput is the input of Parser interface.
type ParseInput[I comparable] interface {
	// Read reads len(slice) and returns the length that the parser could read.
	Read([]I) (int, error)
	// Seek conditions the offset of the input.
	Seek(int, SeekMode) (int, error)
}

func NewParser[I comparable, O ParseOutput](fn ParserFn[I, O]) Parser[I, O] {
	return Parser[I, O]{fn: fn}
}

func (p *Parser[I, O]) Parse(input ParseInput[I]) (ParseInput[I], O, ParseError) {
	return p.fn(input)
}

// SeekMode specifies the behavior of Seek() method.
// It's similar to io.SeekMode.
type SeekMode int

const (
	// SeekModeStart indicates the Seek() method seeks from the start of the input.
	SeekModeStart SeekMode = io.SeekStart
	// SeekModeCurrent indicates the Seek() method seeks from the current offset of the input.
	SeekModeCurrent SeekMode = io.SeekCurrent
	// SeekModeEnd     SeekMode = io.SeekEnd
)

// ParseOutput is the actual type of the parser's output.
// Note that this interface may be constrainted more in future.
type ParseOutput interface{}

// ParseError represents the error of parsers in all parsers.
type ParseError interface {
	error
}

// ErrorIs checks the given error implements ParseError interface.
func ErrorIs[T ParseError](err error, ty T) bool {
	_, ok := err.(T)
	return ok
}

// NoLeftInputToParseError notifies the given input to parser is empty.
type NoLeftInputToParseError struct{}

// Error implements error interface.
func (e *NoLeftInputToParseError) Error() string {
	return "no left input to parse"
}
