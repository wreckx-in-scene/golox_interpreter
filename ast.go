package main

//creating ast structs
type Expr interface{}

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
