package mathskills

import (
	"sort"
)

// steps
// 1.Sort thr data in ascending order
// 2.Find the index of the middle number
// 3.Return number there if the population size is odd
// 4.If even divide the (n + n+1) by 2
func CalcMedian(data []float64) float64 {
	sort.Float64s(data)
	mid := (len(data) / 2)
	if len(data)%2 == 0 {
		return (data[mid] + data[mid] + 1) / 2
	} else {
		return data[mid]
	}
}
