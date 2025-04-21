package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	"strconv" // This package may be helpful...
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	// Use ast.Inspect to traverse the AST
	ast.Inspect(node, func(n ast.Node) bool {
		// Check if the node is a call expression
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true // Continue traversal
		}

		// Check if the call is to expr.ParseAndEval
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true // Continue traversal
		}

		// Check if the selector (X) is an identifier
		xIdent, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return true // Continue traversal
		}

		// Check if the identifier is "expr" and the method is "ParseAndEval"
		if xIdent.Name != "expr" || selectorExpr.Sel.Name != "ParseAndEval" {
			return true // Continue traversal
		}

		// Check if there are exactly two arguments
		if len(callExpr.Args) != 2 {
			return true // Continue traversal
		}

		// Check if the first argument is a string literal
		strLit, ok := callExpr.Args[0].(*ast.BasicLit)
		if !ok || strLit.Kind != token.STRING {
			return true // Continue traversal
		}

		// Get the string value (removing the quotes)
		strValue, err := strconv.Unquote(strLit.Value)
		if err != nil {
			return true // Continue traversal
		}

		// Parse the string into an expr.Expr
		parsed, err := expr.Parse(strValue)
		if err != nil {
			// If there's an error parsing, don't modify the call
			return true
		}

		// Simplify the expression with an empty environment
		simplified := simplify.Simplify(parsed, expr.Env{})

		// Format the simplified expression as a string
		simplifiedStr := expr.Format(simplified)

		// Replace the first argument with the simplified expression
		callExpr.Args[0] = &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(simplifiedStr),
		}

		return true // Continue traversal
	})
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
