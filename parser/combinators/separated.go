package combinators

import (
	"github.com/claudioluciano/gocustomparser/parser"
)

func Separated1[I comparable, EO parser.ParseOutput, SO parser.ParseOutput](element parser.Parser[I, EO], separator parser.Parser[I, SO]) parser.Parser[I, []EO] {
	return parser.NewParser(func(input parser.ParseInput[I]) (parser.ParseInput[I], []EO, parser.ParseError) {
		output := make([]EO, 0)
		rest, e1, err := element.Parse(input)
		if err != nil {
			return rest, output, err
		}

		output = append(output, e1)

		for {
			var eo EO

			// we mustn't generate an error if separator parser fails
			// because such as case is the end of the separated-list.
			rest, _, err = separator.Parse(rest)
			if err != nil {
				break
			}

			rest, eo, err = element.Parse(rest)
			if err != nil {
				// must generate an error.
				return rest, output, err
			}

			output = append(output, eo)
		}

		// we mustn't return err if separator parser fails
		// because such as case is the end of the separated-list.
		return rest, output, nil
	})
}
