package mathskills

import "testing"

// testing for standard deviation function , compare  output to expected input
func TestCalcStdDev(t *testing.T) {
	tests := []struct {
		variance float64
		expected float64
	}{
		{113068.55999999998, 336.2566876658366},
		{20000, 141.4213562373095},
	}

	for _, test := range tests {
		output := CalcStdDev(test.variance)
		if output != test.expected {
			t.Errorf("calculate standard deviation(%v) = %v; Expected %v", test.variance, output, test.expected)
		}
	}
}
