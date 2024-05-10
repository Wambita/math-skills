package mathskills

import (
	"fmt"
	"math"
)

// safely rounds afloat64 to int 64 , checking for potential overflows
// number should not be greater than max int64 value or less than minimum int64 value
// converts to int64 then rounds the number
func SafeRound(num float64) (int64, error) {
	if num > float64(math.MaxInt64) || num < float64(math.MinInt64) {
		return 0, fmt.Errorf("value out of int64 range")
	}
	return int64(math.Round(num)), nil
}
