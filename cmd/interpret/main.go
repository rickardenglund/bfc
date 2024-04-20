package main

import (
	"fmt"
	"os"

	"bfc/gen"
	"bfc/lex"
	"bfc/parse"
)

func main() {
	if len(os.Args) != 2 {
		panic("expected 1 file as argument")
	}

	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens, err := lex.Lex([]rune(string(bs)))
	if err != nil {
		panic(err)
	}

	program, err := parse.Parse(tokens)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Prog: %s\n", program)

	//fmt.Printf("Running\n")
	//runner.Run(program)
	prog := gen.Generate(program)
	fmt.Printf("%s\n", prog)
}
