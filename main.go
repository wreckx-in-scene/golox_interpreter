package main

import "fmt"

func main() {
	source := "var x = 5 + 3;"

	lexer := &Lexer{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}

	tokens := lexer.ScanToken()

	for _, token := range tokens {
		fmt.Println(token.Type, token.Lexeme)
	}
}
