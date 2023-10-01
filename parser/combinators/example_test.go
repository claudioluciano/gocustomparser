package combinators_test

import (
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser"
	"github.com/claudioluciano/gocustomparser/parser/byteparse"
	"github.com/claudioluciano/gocustomparser/parser/combinators"
	"github.com/claudioluciano/gocustomparser/parser/strparse"
)

func ExamplePreceded() {
	predecessor := strparse.Rune('*')
	successor := strparse.Rune('a')
	p := combinators.Preceded(predecessor, successor)

	i := strparse.NewCompleteInput("*a")
	_, o, err := p.Parse(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	// a
	// <nil>
}

func ExampleTerminated() {
	predecessor := strparse.Rune('a')
	successor := strparse.Rune('+')
	p := combinators.Terminated(predecessor, successor)

	i := strparse.NewCompleteInput("a+")
	_, o, err := p.Parse(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	// a
	// <nil>
}

func ExampleSeparated1() {
	element := strparse.Digit1()
	separator := strparse.Rune('|')
	p := combinators.Separated1(element, separator)

	i := strparse.NewCompleteInput("123|456|789Drumato")
	_, o, err := p.Parse(i)
	fmt.Printf("%d\n", len(o))
	fmt.Printf("%s %s %s\n", o[0], o[1], o[2])
	fmt.Println(err)
	// Output:
	// 3
	// 123 456 789
	// <nil>
}

func ExampleSatisfy() {
	i := strparse.NewCompleteInput("abc")
	p := combinators.Satisfy(func(ch rune) bool {
		return ch == 'a'
	})
	_, o, err := p.Parse(i)
	fmt.Printf("%c\n", o)
	fmt.Println(err)
	// Output:
	// a
	// <nil>
}

func ExampleMap() {
	subsubP := strparse.Rune('a')
	subP := combinators.Many1(subsubP)
	p := combinators.Map(subP, func(s []rune) (int, error) { return len(s), nil })

	i := strparse.NewCompleteInput("aaaabaaaa")
	_, o, err := p.Parse(i)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	// 4
	// <nil>
}

func ExampleAlt() {
	p1 := strparse.Rune('a')
	p2 := strparse.Rune('b')
	p := combinators.Many1(combinators.Alt(p1, p2))

	i := strparse.NewCompleteInput("abababc")
	_, o, err := p.Parse(i)
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// ababab
	// <nil>
}

func ExampleDelimited() {
	begin := strparse.Rune('(')
	end := strparse.Rune(')')
	contents := strparse.Digit1()
	p := combinators.Delimited(begin, contents, end)

	i := strparse.NewCompleteInput("(12321)")
	_, o, err := p.Parse(i)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	// 12321
	// <nil>
}

func ExampleMany0() {
	p := combinators.Many0(strparse.Rune('a'))

	i := strparse.NewCompleteInput("baaaa")
	_, o, err := p.Parse(i)
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	//
	// <nil>
}

func ExampleMany1() {
	p := combinators.Many1(strparse.Rune('a'))

	i := strparse.NewCompleteInput("aaaabaa")
	_, o, err := p.Parse(i)
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// aaaa
	// <nil>
}

func ExampleManyMinMax() {
	p := combinators.ManyMinMax(strparse.Rune('a'), 3, 5)

	i := strparse.NewCompleteInput("aaaabbb")
	_, o, err := p.Parse(i)
	fmt.Println(string(o))
	fmt.Println(err)
	// Output:
	// aaaa
	// <nil>
}

func ExampleTake() {
	p := combinators.Take(5, strparse.Rune('a'))

	i := strparse.NewCompleteInput("aaaaabbb")
	_, o, err := p.Parse(i)
	fmt.Println(len(o))
	fmt.Println(err)
	// Output:
	// 5
	// <nil>
}

func ExampleBranches() {
	m := make(map[byte]parser.Parser[byte, string])
	m[0x00] = combinators.Map(byteparse.UInt8(), func(v uint8) (string, error) { return "0x00", nil })
	m[0x01] = combinators.Map(byteparse.UInt8(), func(v uint8) (string, error) { return "0x01", nil })

	p := combinators.Many1(combinators.Branches(m))

	i := byteparse.NewCompleteInput([]byte{0x00, 0x01, 0x00, 0x01, 0x02})
	_, o, err := p.Parse(i)
	fmt.Println(len(o))
	fmt.Printf("%s %s %s %s\n", o[0], o[1], o[2], o[3])
	fmt.Println(err)
	// Output:
	// 4
	// 0x00 0x01 0x00 0x01
	// <nil>
}

func ExampleSequence() {
	a := strparse.Rune('a')
	b := strparse.Rune('b')
	c := strparse.Rune('c')
	p := combinators.Sequence([]parser.Parser[rune, rune]{a, b, c})

	i := strparse.NewCompleteInput("abc")
	_, o, err := p.Parse(i)
	fmt.Println(len(o))
	fmt.Printf("%c %c %c\n", o[0], o[1], o[2])
	fmt.Println(err)
	// Output:
	// 3
	// a b c
	// <nil>
}

func ExampleTwin() {
	one := strparse.Rune('1')
	two := combinators.Map(strparse.Rune('2'), func(s rune) (string, error) {
		return "two", nil
	})
	p := combinators.Twin(one, two)

	i := strparse.NewCompleteInput("12")
	_, o, err := p.Parse(i)
	fmt.Println(string(o.One))
	fmt.Println(o.Two)
	fmt.Println(err)
	// Output:
	// 1
	// two
	// <nil>
}
