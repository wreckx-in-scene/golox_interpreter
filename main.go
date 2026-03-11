package main

import "fmt"

func main() {
	source := "2 + 3 * 4"

	lexer := &Lexer{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}

	tokens := lexer.ScanToken()
	parser := NewParser(tokens)
	ast := parser.Parse()

	result := evaluate(ast)
	fmt.Println(result)
}
