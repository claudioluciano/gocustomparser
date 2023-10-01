package byteparse_test

import (
	"encoding/binary"
	"fmt"

	"github.com/claudioluciano/gocustomparser/parser/byteparse"
)

func ExampleUInt8() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03})
	p := byteparse.UInt8()
	_, o, err := p.Parse(b)
	fmt.Println(o)
	fmt.Println(err)
	// Output:
	//
	// 1
	// <nil>
}

func ExampleUInt16() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03})
	p := byteparse.UInt16(binary.BigEndian)
	_, o, err := p.Parse(b)
	fmt.Printf("0x%x\n", o)
	fmt.Println(err)
	// Output:
	//
	// 0x102
	// <nil>
}

func ExampleUInt32() {
	b := byteparse.NewCompleteInput([]byte{0x01, 0x02, 0x03, 0x04})
	p := byteparse.UInt32(binary.BigEndian)
	_, o, err := p.Parse(b)
	fmt.Printf("0x%x\n", o)
	fmt.Println(err)
	// Output:
	//
	// 0x1020304
	// <nil>
}

func ExampleTag() {
	t := []byte{0x7f, 0x45, 0x4c, 0x46}

	b := byteparse.NewCompleteInput([]byte{0x7f, 0x45, 0x4c, 0x46, 0x02})
	p := byteparse.Tag(t)
	_, o, err := p.Parse(b)
	fmt.Printf("%d\n", len(o))
	fmt.Printf("%x %x %x %x\n", o[0], o[1], o[2], o[3])
	fmt.Println(err)
	// Output:
	// 4
	// 7f 45 4c 46
	// <nil>
}
