package combinators

import "github.com/claudioluciano/gocustomparser/parser"

// Delimited initializes a parser that parses a delimited sequence. (e.g. '[foobar]')
func Delimited[
	I comparable,
	O1 parser.ParseOutput,
	O2 parser.ParseOutput,
	O3 parser.ParseOutput,
](
	begin parser.Parser[I, O1],
	contents parser.Parser[I, O2],
	end parser.Parser[I, O3],
) parser.Parser[I, O2] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], O2, parser.ParseError) {
		var o2 O2
		rest, _, err := begin.Parse(input)
		if err != nil {
			return rest, o2, err
		}

		rest, o2, err = contents.Parse(rest)
		if err != nil {
			return rest, o2, err
		}

		rest, _, err = end.Parse(rest)
		if err != nil {
			return rest, o2, err
		}

		return rest, o2, nil
	})
}
