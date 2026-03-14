golox
A tree-walk interpreter for the Lox programming language, built in Go.
Built from scratch as part of learning systems programming in Go, inspired by Crafting Interpreters by Robert Nystrom.
What is golox?
golox is a complete interpreter for the Lox language — a dynamically typed, garbage-collected scripting language designed by Robert Nystrom for his book Crafting Interpreters.
This implementation follows the tree-walk interpreter approach: source code is scanned into tokens, parsed into an Abstract Syntax Tree (AST), and then evaluated by walking the tree recursively.

How it works
Source code flows through three stages before producing output:

Stage	File	Input	Output
Lexer / Scanner	lexer.go, token.go	Raw source string	[]Token slice
Parser	parser.go, ast.go	[]Token slice	AST (nested structs)
Interpreter	interpreter.go	AST root node	Evaluated result

Stage 1 — Lexer
The lexer scans source code character by character, grouping characters into meaningful tokens. Each token has a type, lexeme (raw text), optional literal value, and line number for error reporting.
Handles: single-character tokens, one-or-two character tokens (!=, ==, <=, >=), string literals, number literals, identifiers, and 16 reserved keywords.
Stage 2 — Parser
The parser uses recursive descent — a top-down parsing technique where each grammar rule maps to a function. The call chain naturally encodes operator precedence:

expression() → assignment() → equality() → comparison()
    → term() → factor() → unary() → primary()

Lower in the chain = higher precedence = evaluated first. This is how 2 + 3 * 4 correctly evaluates to 14, not 20.
Stage 3 — Interpreter
The interpreter walks the AST recursively using Go's type switch. Each node type has a corresponding evaluation rule:

evaluate(Binary{+, Literal{2}, Binary{*, Literal{3}, Literal{4}}})
  → left  = evaluate(Literal{2})              = 2
  → right = evaluate(Binary{*, Literal{3}, Literal{4}})
              → 3 * 4                          = 12
  → 2 + 12                                     = 14


Features
Variables and scoping
Variables are stored in an Environment — a hashmap of name → value. Environments chain together for lexical scoping: each block creates a new environment that points to its enclosing one.

var x = 10;
var name = "Tamaghna";
print x + 5;   // 15

Control flow
Full if/else, while loops, and for loops. For loops are desugared into while loops at parse time — no new AST nodes needed.

for (var i = 0; i < 5; i = i + 1) {
    print i;
}

if (x > 10) {
    print "big";
} else {
    print "small";
}

Functions and recursion
First-class functions with lexical scoping. Return values use Go's panic/recover mechanism for clean unwinding of the call stack.

fun fibonacci(n) {
    if (n <= 1) { return n; }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

print fibonacci(10);   // 55

Expressions
•Arithmetic: +, -, *, /
•Comparison: >, >=, <, <=
•Equality: ==, !=
•Logical: and, or, !
•String concatenation via +
•Grouping with parentheses

Project structure
File	Responsibility
token.go	TokenType constants and Token struct definition
lexer.go	Scanner — converts source string into token slice
ast.go	AST node definitions (Expr and Stmt interfaces + all node types)
parser.go	Recursive descent parser — converts tokens into AST
environment.go	Variable storage with lexical scoping chain
interpreter.go	Tree-walk evaluator and statement executor
main.go	Entry point — wires lexer → parser → interpreter

Running golox
Prerequisites
•Go 1.22 or later
•Zero external dependencies

Clone and build

git clone https://github.com/yourusername/golox
cd golox
go build -o golox.exe .    # Windows
go build -o golox .        # Mac / Linux

Mode 1 — Run a .lox file
Write your Lox code in a file and pass it as an argument:

# create a file
# hello.lox:
#   var name = "world";
#   print "Hello " + name;

.\golox.exe hello.lox     # Windows
./golox hello.lox          # Mac / Linux

Mode 2 — Interactive REPL
Run without any arguments to enter the interactive shell. Type Lox expressions line by line and see results instantly:

.\golox.exe              # Windows
./golox                   # Mac / Linux

> var x = 10;
> print x * 2;
20
> fun add(a, b) { return a + b; }
> print add(3, 4);
7
> exit

Example .lox program
Save this as fibonacci.lox and run it:

fun fibonacci(n) {
    if (n <= 1) { return n; }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

for (var i = 0; i < 10; i = i + 1) {
    print fibonacci(i);
}


What I learned
This project was built as part of learning Go — started with zero Go knowledge and built up through two prior projects (a CLI todo app and a Snake game with goroutines).

Go concepts applied
•Structs and methods — lexer, parser, interpreter, environment
•Interfaces — Expr and Stmt as empty interfaces with type switches
•Slices and maps — token list, keyword map, environment values
•Error handling patterns — Go's idiomatic err return style
•Panic/recover — used for return statement unwinding
•Multiple return values — used throughout for value + error pairs

Built with Go  •  Inspired by Crafting Interpreters by Robert Nystr
