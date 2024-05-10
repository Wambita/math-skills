package mathskills

import "testing"

func TestCalcAverage(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{input: []float64{101, 242, 999, 98, 456}, expected: 379.2},
		{input: []float64{100, 200, 500, 300, 400}, expected: 300},
	}

	for _, test := range tests {
		output := CalcAverage(test.input)
		if output != test.expected {
			t.Errorf("calculateAverage(%v) = %v; Expected %v", test.input, output, test.expected)
		}
	}
}
