package mathskills
func CalcVariance(data []float64, average float64) float64{
	sum := 0.0

	for _, num := range data {
		num = (num - average) * (num - average)
		sum += num
	}
	return  sum / float64(len(data))

}