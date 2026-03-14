package main

func main() {
	source := `
var a = 0;
var b = 1;
var n = 6; // Calculate 6th Fibonacci number roughly

while (n > 0) {
    print a;
    var temp = a;
    a = b;
    b = temp + b;
    n = n - 1;
}

if (a == 8) {
    print "Fibonacci logic: Success";
} else {
    print "Fibonacci logic: Failed";
}
`

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
