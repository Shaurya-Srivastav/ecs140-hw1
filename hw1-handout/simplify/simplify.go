package simplify

import (
	"hw1/expr"
)

// Simplify should return the simplified expression
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	switch e := e.(type) {
	case expr.Literal:
		// Literals are already in their simplest form
		return e

	case expr.Var:
		// If the variable exists in the environment, replace it with its value
		if val, ok := env[e]; ok {
			return expr.Literal(val)
		}
		// Otherwise keep it as a variable
		return e

	case expr.Unary:
		// Simplify the operand
		x := Simplify(e.X, env)

		// If the operand is a literal, compute the result
		if lit, ok := x.(expr.Literal); ok {
			switch e.Op {
			case '+':
				return lit
			case '-':
				return expr.Literal(-float64(lit))
			}
		}

		// Otherwise, keep the unary operator with the simplified operand
		return expr.Unary{Op: e.Op, X: x}

	case expr.Binary:
		// Simplify both operands
		x := Simplify(e.X, env)
		y := Simplify(e.Y, env)

		// If both operands are literals, compute the result
		xLit, xIsLit := x.(expr.Literal)
		yLit, yIsLit := y.(expr.Literal)

		if xIsLit && yIsLit {
			switch e.Op {
			case '+':
				return expr.Literal(float64(xLit) + float64(yLit))
			case '-':
				return expr.Literal(float64(xLit) - float64(yLit))
			case '*':
				return expr.Literal(float64(xLit) * float64(yLit))
			case '/':
				return expr.Literal(float64(xLit) / float64(yLit))
			}
		}

		// Special case: addition with 0
		if xIsLit && float64(xLit) == 0 && e.Op == '+' {
			return y
		}
		if yIsLit && float64(yLit) == 0 && e.Op == '+' {
			return x
		}

		// Special case: multiplication with 0 or 1
		if xIsLit && e.Op == '*' {
			if float64(xLit) == 0 {
				return expr.Literal(0)
			}
			if float64(xLit) == 1 {
				return y
			}
		}
		if yIsLit && e.Op == '*' {
			if float64(yLit) == 0 {
				return expr.Literal(0)
			}
			if float64(yLit) == 1 {
				return x
			}
		}

		// Otherwise, create a new binary expression with simplified operands
		return expr.Binary{Op: e.Op, X: x, Y: y}

	default:
		// If not one of the known types, panic
		panic("unknown Expr type")
	}
}
