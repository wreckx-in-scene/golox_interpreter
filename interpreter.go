package main

import "fmt"

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
			//could be numbers or strings
			l, lok := left.(float64)
			r, rok := right.(float64)

			if lok && rok {
				return l + r
			}

			return fmt.Sprintf("%v%v , left , right")

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
