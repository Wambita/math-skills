package mathskills

import (
	"fmt"
	"math"
)

// safely rounds afloat64 to int 64 , checking for potential overflowss
func SafeRound(num float64) (int64, error) {
	if num > float64(math.MaxInt64) || num < float64(math.MinInt64) {
		return 0, fmt.Errorf("value out of int64 range")
	}
	return int64(math.Round(num)), nil
}
