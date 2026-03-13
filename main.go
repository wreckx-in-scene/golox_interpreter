package main

func main() {
	source := `var name = "Tamaghna";
var age = 20;
print name;
print age + 5;`

	lexer := &Lexer{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}

	tokens := lexer.ScanToken()
	parser := NewParser(tokens)
	statements := parser.Parse()

	for _, stmt := range statements {
		execute(stmt)
	}
}
