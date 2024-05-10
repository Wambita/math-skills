package main

import (
	"fmt"

	mathskills "mathskills/programfiles"
)

func main() {
	data, err := mathskills.OpenFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	// average , kept as a float for accurate calculations
	average := mathskills.CalcAverage(data)

	// calculate median and round it safely checking for overflows
	median, err := mathskills.SafeRound(mathskills.CalcMedian(data))
	if err != nil {
		fmt.Println(err)
	}

	// calculate variance and keep it as float for accurate calculations
	variance := mathskills.CalcVariance(data, average)

	// claculate median and round it safely
	standarddeviation, err := mathskills.SafeRound(mathskills.CalcStdDev(variance))
	if err != nil {
		fmt.Println(err)
	}

	// round the average safely
	roundedAverage, err := mathskills.SafeRound(average)
	if err != nil {
		fmt.Println(err)
	}
	// round the variance safely
	roundedVariance, err := mathskills.SafeRound(variance)
	if err != nil {
		fmt.Println(err)
	}

	// Print all values to the nearest integer
	fmt.Printf("Average: %d\n", roundedAverage)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", roundedVariance)
	fmt.Printf("Standard Deviation: %d\n", standarddeviation)
}
