package mathskills

import (
	"sort"
)

func CalcMedian(data []float64) float64 {
	sort.Float64s(data)
	mid := (len(data) / 2)
	if len(data)%2 == 0 {
		return (data[mid] + data[mid] + 1) / 2
	} else {
		return data[mid]
	}
}
