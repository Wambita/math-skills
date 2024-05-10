package mathskills

// steps
// 1.Declare sum variable
// 2.Loop throught the data and sum each value
// 3.Divide the sum by the population size
func CalcAverage(data []float64) float64 {
	sum := 0.0
	for _, i := range data {
		sum += i
	}
	return sum / float64(len(data))
}
