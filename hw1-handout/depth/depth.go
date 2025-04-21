package depth

import (
	"hw1/expr"
)

// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	switch e := e.(type) {
	case expr.Literal, expr.Var:
		// Base case: Literals and variables are leaf nodes, so depth is 1
		return 1
	case expr.Unary:
		// For unary expressions, count the operator plus max depth of operand
		return 1 + Depth(e.X)
	case expr.Binary:
		// For binary expressions, count the operator plus max depth of its operands
		leftDepth := Depth(e.X)
		rightDepth := Depth(e.Y)

		// Return the max depth
		if leftDepth > rightDepth {
			return 1 + leftDepth
		}
		return 1 + rightDepth
	default:
		// If not one of the known types, panic as shown in the test case
		panic("unknown Expr type")
	}
}
