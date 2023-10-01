package main

import (
	"fmt"
	"os"

	"github.com/claudioluciano/gocustomparser/parser"
)

// const s = `["foo","bar","baz"]`

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: ./jsonparser <json-file>")
		os.Exit(1)
	}

	f, err := os.Open(args[1])
	// f, err := os.Open("./example.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}

	r := parser.NewIOReadSeeker(f)
	p := parseJSONValue()
	_, v, err := p.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(v)
}
