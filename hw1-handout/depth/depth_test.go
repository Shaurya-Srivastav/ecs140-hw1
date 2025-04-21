package depth

import (
	"hw1/expr"
	"testing"
)

// !+Depth
func TestDepth(t *testing.T) {
	tests := []struct {
		expr     string
		expected uint
	}{
		{"10", 1},
		{"xyz", 1},
		{"-10", 2},
		{"-x", 2},
		{"1-x", 2},
		{"1+x", 2},
		{"1-x", 2},
		{"(1 + 2) + (3 + X)", 3},
		{"1 + 2 + 3 + X", 4},
		// Additional test cases for full coverage
		{"1 * 2 + 3 * 4", 3},         // Test binary with depth on left
		{"1 + 2 * (3 + 4 * 5)", 5},   // Corrected expected depth
		{"-(-(-x))", 4},              // Test nested unary operators
		{"5 / (2 - 3 * (1 + 0))", 5}, // Corrected expected depth
	}

	for _, test := range tests {
		e, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// Run the method
		actual := Depth(e)

		// Check the result
		if actual != test.expected {
			t.Errorf("Depth(%s) = %d, expected %d\n", test.expr, actual, test.expected)
		}
	}
}

type Foo int

func (Foo) Eval(env expr.Env) float64 {
	return 0.0
}

func (Foo) Check(vars map[expr.Var]bool) error {
	return nil
}

func TestDepth_Fail(t *testing.T) {

	func() {
		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()

		var f Foo
		Depth(f)
	}()

}

//!-Depth
