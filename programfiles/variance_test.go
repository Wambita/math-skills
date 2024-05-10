package mathskills

import "testing"

// testing for variance function , compare  output to expected input
func TestCalcVariance(t *testing.T) {
	tests := []struct {
		input    []float64
		average  float64
		expected float64
	}{
		{input: []float64{101, 242, 999, 98, 456}, average: 379.2, expected: 113068.55999999997},
		{input: []float64{100, 200, 500, 300, 400}, average: 300, expected: 20000},
	}

	for _, test := range tests {
		output := CalcVariance(test.input, test.average)
		if output != test.expected {
			t.Errorf("calculateVariance(%v) = %v; Expected %v", test.input, output, test.expected)
		}
	}
}
