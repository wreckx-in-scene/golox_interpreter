package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run(source string) {
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

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: err")
		os.Exit(1)
	}

	run(string(data))
}

func runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("golox REPL - type 'exit' to quit")
	for {
		fmt.Print(">")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "exit" {
			break
		}
		if line == "" {
			continue
		}

		run(line)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		runREPL()
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		fmt.Println("Usage: golox [script.lox]")
		os.Exit(1)
	}
}
