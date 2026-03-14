package main

//creating ast structs
type Expr interface{}

//base interface for statements
type Stmt interface{}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

type Unary struct {
	Operator Token
	Right    Expr
}

type Literal struct {
	Value interface{}
}

type Grouping struct {
	Expression Expr
}

//adding stmt nodes

//print statement
type PrintStmt struct {
	Expression Expr
}

//var statement
type VarStmt struct {
	Name        Token
	Initializer Expr
}

// expression used as a statement
type ExprStmt struct {
	Expression Expr
}

//identifier node
type Identifier struct {
	Name Token
}

//if statement
type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

//while statement
type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

type BlockStmt struct {
	Statements []Stmt
}

type AssignStmt struct {
	Name  Token
	Value Expr
}
