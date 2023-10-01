package strparse_test

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser/strparse"
)

func ExampleRune() {
	i := strparse.NewCompleteInput("abc")
	p := strparse.Rune('a')
	_, o, err := p.Parse(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	//
	// a
	// <nil>
}

func ExampleTag() {
	i := strparse.NewCompleteInput("Drumato")
	p := strparse.Tag("Drum")
	_, o, err := p.Parse(i)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// Drum
	// <nil>
}

func ExampleDigit1() {
	i := strparse.NewCompleteInput("112233abc")
	p := strparse.Digit1()
	_, o, err := p.Parse(i)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// 112233
	// <nil>
}
