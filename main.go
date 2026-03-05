package main

import "fmt"

func main() {
	source := `var x = 42;
	var name = "Amogh";
	if (x > 10) {
		print name;
	}`

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
