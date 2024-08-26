package main

import (
	"os"
	// "github.com/go-parser/src/lexer"
	"github.com/go-parser/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/03.lang")

	// tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(string(bytes))
	litter.Dump(ast)

}
