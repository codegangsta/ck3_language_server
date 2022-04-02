package main

import (
	"fmt"

	"github.com/codegangsta/ck3_language_server/parser"
)

func main() {
	ast, _ := parser.ParseFile("parser/test_cases/example.txt")
	fmt.Printf("%#v", ast)
}
