package main

import (
	"fmt"
	"os"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
)

type myTokenKind string

const (
	myTokenKindInteger myTokenKind = "<Integer>"
	myTokenKindLParen  myTokenKind = "("
	myTokenKindRParen  myTokenKind = ")"
)

type myToken struct {
	kind  myTokenKind
	p     position
	value int
}

// String() implemlents fmt.Stringer interface
func (t myToken) String() string {
	return fmt.Sprintf("(%d:%d) kind: %+v, value: %d", t.p.line, t.p.column, t.kind, t.value)
}

type position struct {
	line   uint
	column uint
}

type pseudoTokenStream struct {
	tokens []myToken
	offset int
}

func (c *pseudoTokenStream) Read(buf []myToken) (int, error) {
	if c.offset >= len(c.tokens) {
		return 0, &parser.NoLeftInputToParseError{}
	}

	copy(buf, c.tokens[c.offset:])
	c.offset += len(buf)

	return len(buf), nil
}

func (c *pseudoTokenStream) Seek(n int, mode parser.SeekMode) (int, error) {
	switch mode {
	case parser.SeekModeStart:
		if n >= len(c.tokens) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset = n
		return c.offset, nil
	case parser.SeekModeCurrent:
		if c.offset+n >= len(c.tokens) {
			return 0, &parser.NoLeftInputToParseError{}
		}

		c.offset += n
		return c.offset, nil
	default:
		panic("given seek mode is not supported")
	}
}

func main() {
	p := setupParser()
	succeedCase(p)
	failCase(p)
}

func succeedCase(p parser.Parser[myToken, int]) {
	tokens := []myToken{
		// "("
		{kind: myTokenKindLParen, p: position{line: 1, column: 1}},
		// 12345
		{kind: myTokenKindInteger, p: position{line: 1, column: 2}, value: 12345},
		// ")"
		{kind: myTokenKindRParen, p: position{line: 1, column: 7}},
	}
	s := &pseudoTokenStream{tokens: tokens}

	_, v, err := p.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(v)
}

func failCase(p parser.Parser[myToken, int]) {
	tokens := []myToken{
		// "("
		{kind: myTokenKindLParen, p: position{line: 1, column: 1}},
		// 12345
		{kind: myTokenKindInteger, p: position{line: 1, column: 2}, value: 12345},
		// 678910
		{kind: myTokenKindInteger, p: position{line: 1, column: 7}, value: 678910},
	}
	s := &pseudoTokenStream{tokens: tokens}

	_, _, err := p.Parse(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	}
}

// parser := "(" integer-token ")"
func setupParser() parser.Parser[myToken, int] {
	// begin := "("
	begin := combinators.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindLParen
	})

	// contents := integer-token
	contents := combinators.Map(combinators.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindInteger
	}), func(t myToken) (int, error) {
		return t.value, nil
	})

	// end := ")"
	end := combinators.Satisfy(func(t myToken) bool {
		return t.kind == myTokenKindRParen
	})
	return combinators.Delimited(begin, contents, end)
}
