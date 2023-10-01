package combinators

import "github.com/claudioluciano/gocustomparser/parser"

func Lazy[I comparable, O parser.ParseOutput](fn func() parser.Parser[I, O]) parser.Parser[I, O] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O, parser.ParseError) {
		p := fn()
		return p.Parse(input)
	})
}
