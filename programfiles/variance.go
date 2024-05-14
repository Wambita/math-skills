package mathskills

import "math"

// steps to calculate variance
// 1. Find the square of the difference of each number and the average
// 2. sum the squares
// 3. divide the sum with the population size
func CalcVariance(data []float64, average float64) float64 {
	sum := 0.0

	for _, num := range data {
		num = math.Pow((num - average), 2)
		sum += num
	}
	return sum / float64(len(data))
}
