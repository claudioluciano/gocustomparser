package lexer

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

type CustomLexemeFuncOptions struct {
	Value       []byte
	Startline   int
	Endline     int
	Startcolumn int
	Endcolumn   int
	Scan        *lexmachine.Scanner
}

type CustomLexeme struct {
	Regex []byte                                              `validate:"required"`
	Func  func(opts *CustomLexemeFuncOptions) (*Token, error) `validate:"required"`
}

type Lexer struct {
	lexmachine *lexmachine.Lexer
	literals   []TokenKind
	keywords   []TokenKind
	customs    []*CustomLexeme
}

type LexerOptions struct {
	Literals []TokenKind
	Keywords []TokenKind `validate:"required"`
	Customs  []*CustomLexeme
}

// New creates a new Lexer
// The lexer will not match anything unless you add a rule for it.
func NewLexer(opts *LexerOptions) (*Lexer, error) {
	if err := validator.New().Struct(opts); err != nil {
		return nil, err
	}

	l := &Lexer{
		literals:   opts.Literals,
		keywords:   opts.Keywords,
		customs:    opts.Customs,
		lexmachine: lexmachine.NewLexer(),
	}

	if err := l.compileLexer(); err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Lexer) Lex(input string) ([]*Token, error) {
	scanner, err := l.lexmachine.Scanner([]byte(input))
	if err != nil {
		return nil, err
	}

	var tokens []*Token
	for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, tok.(*Token))
	}

	return tokens, nil
}

func (l *Lexer) compileLexer() error {
	for _, lit := range l.literals {
		r := "\\" + strings.Join(strings.Split(string(lit), ""), "\\")
		l.lexmachine.Add([]byte(r), token(lit))
	}

	for _, name := range l.keywords {
		l.lexmachine.Add([]byte(strings.ToLower(string(name))), token(name))
		l.lexmachine.Add([]byte(strings.ToUpper(string(name))), token(name))
	}

	for _, sk := range l.customs {
		l.lexmachine.Add(sk.Regex, rule(sk))
	}

	if err := l.lexmachine.Compile(); err != nil {
		return err
	}

	return nil
}

func rule(sk *CustomLexeme) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		// if the function is nil, the token will be skipped
		if sk.Func == nil {
			return nil, nil
		}

		return sk.Func(&CustomLexemeFuncOptions{
			Value:       m.Bytes,
			Startline:   m.StartLine,
			Endline:     m.EndLine,
			Startcolumn: m.StartColumn,
			Endcolumn:   m.EndColumn,
			Scan:        s,
		})
	}
}

func token(kind TokenKind) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return NewToken(&TokenOptions{
			Kind:        kind,
			Value:       string(m.Bytes),
			StartLine:   m.StartLine,
			StartColumn: m.StartColumn,
			EndLine:     m.EndLine,
			EndColumn:   m.EndColumn,
		}), nil
	}
}
