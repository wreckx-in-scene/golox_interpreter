package main

import "fmt"

var env = NewEnvironment()

//return value wrapper
type ReturnValue struct {
	Value interface{}
}

//lox function
type LoxFunction struct {
	Declaration FunctionStmt
}

func (f LoxFunction) call(args []interface{}) (result interface{}) {
	funcEnv := NewEnvironment()
	funcEnv.enclosing = env

	for i, param := range f.Declaration.Params {
		funcEnv.define(param.Lexeme, args[i])
	}

	previous := env
	env = funcEnv

	defer func() {
		env = previous
		if r := recover(); r != nil {
			if ret, ok := r.(ReturnValue); ok {
				result = ret.Value
			}
		}
	}()

	for _, stmt := range f.Declaration.Body {
		execute(stmt)
	}

	return nil
}

func evaluate(expr Expr) interface{} {
	switch e := expr.(type) {

	case Literal:
		return e.Value
	case Grouping:
		return evaluate(e.Expression)
	case Unary:
		right := evaluate(e.Right)
		switch e.Operator.Type {
		case MINUS:
			return -right.(float64)
		case BANG:
			return !isTruthy(right)
		}

	case Binary:
		left := evaluate(e.Left)
		right := evaluate(e.Right)

		switch e.Operator.Type {
		case PLUS:
			l, lok := left.(float64)
			r, rok := right.(float64)
			if lok && rok {
				return l + r
			}
			return fmt.Sprintf("%v%v", left, right)
		case MINUS:
			return left.(float64) - right.(float64)
		case STAR:
			return left.(float64) * right.(float64)
		case SLASH:
			return left.(float64) / right.(float64)
		case GREATER:
			return left.(float64) > right.(float64)
		case GREATER_EQUAL:
			return left.(float64) >= right.(float64)
		case LESS:
			return left.(float64) < right.(float64)
		case LESS_EQUAL:
			return left.(float64) <= right.(float64)
		case EQUAL_EQUAL:
			return left == right
		case BANG_EQUAL:
			return left != right
		}

	case Identifier:
		return env.get(e.Name.Lexeme)

	// FIX: Add AssignStmt to evaluate so nested assignments and ExprStmts work
	case AssignStmt:
		value := evaluate(e.Value)
		env.assign(e.Name.Lexeme, value)
		return value

	case CallExpr:
		callee := evaluate(e.Callee)
		var args []interface{}
		for _, arg := range e.Arguments {
			args = append(args, evaluate(arg))
		}
		if fn, ok := callee.(LoxFunction); ok {
			return fn.call(args)
		}
		fmt.Println("Not a function")

	}

	return nil
}

func isTruthy(val interface{}) bool {
	if val == nil {
		return false
	}

	if b, ok := val.(bool); ok {
		return b
	}

	return true
}

//execution statement
func execute(stmt Stmt) {
	switch s := stmt.(type) {
	case PrintStmt:
		value := evaluate(s.Expression)
		fmt.Println(value)

	case VarStmt:
		var value interface{}
		if s.Initializer != nil {
			value = evaluate(s.Initializer)
		}
		env.define(s.Name.Lexeme, value)

	case ExprStmt:
		evaluate(s.Expression)

	// FIX: Use env.assign to update the value, not redefine it
	case AssignStmt:
		value := evaluate(s.Value)
		env.assign(s.Name.Lexeme, value)

	case BlockStmt:
		for _, stmt := range s.Statements {
			execute(stmt)
		}

	case IfStmt:
		condition := evaluate(s.Condition)
		if isTruthy(condition) {
			execute(s.ThenBranch)
		} else if s.ElseBranch != nil {
			execute(s.ElseBranch)
		}

	case WhileStmt:
		for isTruthy(evaluate(s.Condition)) {
			execute(s.Body)
		}
	case FunctionStmt:
		function := LoxFunction{Declaration: s}
		env.define(s.Name.Lexeme, function)

	case ReturnStmt:
		var value interface{}
		if s.Value != nil {
			value = evaluate(s.Value)
		}
		panic(ReturnValue{Value: value})
	}

}
