package min

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{
		in:  []int{-1, 0, 1, 2, 4},
		out: -1,
	},
	{
		in:  []int{1},
		out: 1,
	},
	{
		in:  []int{},
		out: 0,
	},
	{
		in:  nil,
		out: 0,
	},
	// Additional test cases for full coverage
	{
		in:  []int{5, 4, 3, 2, 1},
		out: 1,
	},
	{
		in:  []int{0, -10, 10},
		out: -10,
	},
	{
		in:  []int{5, 5, 5},
		out: 5,
	},
}

func TestMin(t *testing.T) {
	for i, test := range tests {
		m := Min(test.in)
		if m != test.out {
			t.Errorf("#%d: Min(%v)=%d; want %d", i, test.in, m, test.out)
		}
	}
}
