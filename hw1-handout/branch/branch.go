package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// branchCount should count the number of branching statements in the function.
func branchCount(fn *ast.FuncDecl) uint {
	// Counter for branching statements
	count := uint(0)

	// Use ast.Inspect to traverse the function's AST
	ast.Inspect(fn, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.IfStmt:
			// Count if statements
			count++
		case *ast.ForStmt:
			// Count for statements
			count++
		case *ast.RangeStmt:
			// Count range loops
			count++
		case *ast.SwitchStmt:
			// Count switch statements
			count++
		case *ast.TypeSwitchStmt:
			// Count type switch statements
			count++
		}
		return true // Continue traversal
	})

	return count
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
