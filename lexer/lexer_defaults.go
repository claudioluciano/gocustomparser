package lexer

func DefaultIdentifier() []*CustomLexeme {
	return []*CustomLexeme{
		{
			Regex: []byte(`"[a-zA-Z-_][a-zA-Z0-9-_]*"`),
			Func: func(opts *CustomLexemeFuncOptions) (*Token, error) {
				value := opts.Value
				// remove the quotes
				value = value[1 : len(value)-1]

				return NewToken(&TokenOptions{
					Kind:        TokenKindIdentifier,
					Value:       string(value),
					StartLine:   opts.Startline,
					StartColumn: opts.Startcolumn,
					EndLine:     opts.Endline,
					EndColumn:   opts.Endcolumn,
				}), nil
			},
		},
		{
			Regex: []byte(`[a-zA-Z-_][a-zA-Z0-9-_]*`),
			Func: func(opts *CustomLexemeFuncOptions) (*Token, error) {
				return NewToken(&TokenOptions{
					Kind:        TokenKindIdentifier,
					Value:       string(opts.Value),
					StartLine:   opts.Startline,
					StartColumn: opts.Startcolumn,
					EndLine:     opts.Endline,
					EndColumn:   opts.Endcolumn,
				}), nil
			},
		},
	}
}

func DefaultNumber() []*CustomLexeme {
	return []*CustomLexeme{
		{
			Regex: []byte(`[0-9]*\.?[0-9]+`),
			Func: func(opts *CustomLexemeFuncOptions) (*Token, error) {
				return NewToken(&TokenOptions{
					Kind:        TokenKindNumber,
					Value:       string(opts.Value),
					StartLine:   opts.Startline,
					StartColumn: opts.Startcolumn,
					EndLine:     opts.Endline,
					EndColumn:   opts.Endcolumn,
				}), nil
			},
		},
	}
}

func DefaultWhitespace() []*CustomLexeme {
	return []*CustomLexeme{
		{
			Regex: []byte(`( |\t|\n|\r)+`),
			Func:  nil,
		},
	}
}
