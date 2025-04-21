package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{30001, 6, 2, UnknownTriangle},
		// Cases where values exceed maximum thresholds
		{5, 20001, 7, UnknownTriangle},
		{5, 7, 10001, UnknownTriangle},

		// Cases with invalid values (zero or negative)
		{0, 5, 7, UnknownTriangle},
		{5, 0, 7, UnknownTriangle},
		{5, 7, 0, UnknownTriangle},
		{-1, 5, 7, UnknownTriangle},
		{5, -1, 7, UnknownTriangle},
		{5, 7, -1, UnknownTriangle},

		// Invalid triangles (sum of two sides <= third side)
		{1, 2, 3, InvalidTriangle},
		{3, 1, 2, InvalidTriangle},
		{2, 3, 1, InvalidTriangle},

		// Right triangles (a² = b² + c²)
		{5, 3, 4, RightTriangle},
		{13, 5, 12, RightTriangle},

		// Acute triangles (a² < b² + c²)
		{4, 4, 4, AcuteTriangle},
		{5, 6, 7, AcuteTriangle},

		// Obtuse triangles (a² > b² + c²)
		{10, 5, 6, ObtuseTriangle},
		{9, 5, 5, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
