package main

import "fmt"

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// peek function
func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

// advance function
func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.tokens[p.current-1]
}

// checking tokenType
func (p *Parser) check(tokType TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.tokens[p.current].Type == tokType
}

// matching function
func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

//parsing functions

// 1 - primary function
func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return Literal{Value: false}
	}

	if p.match(TRUE) {
		return Literal{Value: true}
	}

	if p.match(NIL) {
		return Literal{Value: nil}
	}

	if p.match(NUMBER, STRING) {
		return Literal{Value: p.tokens[p.current-1].Literal}
	}

	if p.match(LEFT_PAREN) {
		expr := p.expression()
		p.match(RIGHT_PAREN)
		return Grouping{Expression: expr}
	}

	if p.match(IDENTIFIER) {
		return Identifier{Name: p.tokens[p.current-1]}
	}

	return nil
}

// unary function
func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.tokens[p.current-1]
		right := p.unary()
		return Unary{Operator: operator, Right: right}
	}
	return p.primary()
}

// factor function
func (p *Parser) factor() Expr {
	left := p.unary()

	for p.match(STAR, SLASH) {
		operator := p.tokens[p.current-1]
		right := p.unary()
		left = Binary{Left: left, Operator: operator, Right: right}
	}

	return left
}

func (p *Parser) term() Expr {
	left := p.factor()

	for p.match(PLUS, MINUS) {
		operator := p.tokens[p.current-1]
		right := p.factor()
		left = Binary{Left: left, Operator: operator, Right: right}
	}

	return left
}

// comparision functions
func (p *Parser) comparision() Expr {
	left := p.term()

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.tokens[p.current-1]
		right := p.term()
		left = Binary{Left: left, Operator: operator, Right: right}
	}

	return left
}

func (p *Parser) equality() Expr {
	left := p.comparision()

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.tokens[p.current-1]
		right := p.comparision()
		left = Binary{Left: left, Operator: operator, Right: right}
	}

	return left
}

func (p *Parser) expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.equality()

	if p.match(EQUAL) {
		equals := p.tokens[p.current-1]
		value := p.assignment()

		if name, ok := expr.(Identifier); ok {
			return AssignStmt{Name: name.Name, Value: value}
		}

		fmt.Println("Error : Inavlid assignment target on line", equals.Line)
	}

	return expr
}

// parse function
func (p *Parser) Parse() []Stmt {
	var statements []Stmt
	for !p.isAtEnd() {
		stmt := p.statement()
		statements = append(statements, stmt)
	}
	return statements
}
func (p *Parser) statement() Stmt {

	if p.match(IF) {
		return p.ifStatement()
	}
	if p.match(WHILE) {
		return p.whileStatement()
	}
	if p.match(LEFT_BRACE) {
		return p.blockStatement()
	}
	if p.match(PRINT) {
		return p.printStatement()
	}
	if p.match(VAR) {
		return p.varStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() Stmt {
	value := p.expression()
	p.match(SEMICOLON)
	return PrintStmt{Expression: value}
}

func (p *Parser) varStatement() Stmt {
	name := p.advance()
	var initializer Expr
	if p.match(EQUAL) {
		initializer = p.expression()
	}

	p.match(SEMICOLON)
	return VarStmt{Name: name, Initializer: initializer}
}

func (p *Parser) expressionStatement() Stmt {
	expr := p.expression()
	p.match(SEMICOLON)
	return ExprStmt{Expression: expr}
}

func (p *Parser) blockStatement() Stmt {
	var statements []Stmt
	for !p.check(RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.statement())
	}
	p.match(RIGHT_BRACE)
	return BlockStmt{Statements: statements}
}

func (p *Parser) ifStatement() Stmt {
	p.match(LEFT_PAREN)
	condition := p.expression()
	p.match(RIGHT_PAREN)

	thenBranch := p.statement()
	var elseBranch Stmt
	if p.match(ELSE) {
		elseBranch = p.statement()
	}

	return IfStmt{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (p *Parser) whileStatement() Stmt {
	p.match(LEFT_PAREN)
	condition := p.expression()
	p.match(RIGHT_PAREN)
	body := p.statement()
	return WhileStmt{Condition: condition, Body: body}
}
