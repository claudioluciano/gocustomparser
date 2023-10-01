package lexer

import "fmt"

type TokenKind string

const (
	TokenKindIdentifier TokenKind = "IDENTIFIER"
	TokenKindNumber     TokenKind = "NUMBER"
)

type Token struct {
	Lexeme      []byte
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
	Kind        TokenKind
	Value       interface{}
}

type TokenOptions struct {
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
	Kind        TokenKind
	Value       interface{}
}

func NewToken(opts *TokenOptions) *Token {
	return &Token{
		StartLine:   opts.StartLine,
		StartColumn: opts.StartColumn,
		EndLine:     opts.EndLine,
		EndColumn:   opts.EndColumn,
		Value:       opts.Value,
		Kind:        opts.Kind,
	}
}

func (t TokenKind) String() string {
	return string(t)
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %q (%d, %d)-(%d, %d)", t.Kind, t.Value, t.StartLine, t.StartColumn, t.EndLine, t.EndColumn)
}
