package main

import (
	"fmt"
	"os"
	"path/filepath"

	mathskills "mathskills/programfiles"
)

func main() {
	// check length og arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . data.txt")
		return
	}
	filePath := os.Args[1]

	// open  data file
	data, err := mathskills.OpenFile(filePath)
	if filepath.Ext(filePath) != ".txt" {
		fmt.Println("wrong file extension: Usage go run . data.txt")
		return
	}

	if filePath != "data.txt" {
		fmt.Println("wrong file name: Usage go run . data.txt")
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// average , kept as a float for accurate calculations
	average := mathskills.CalcAverage(data)

	// calculate median and round it safely checking for overflows
	median, err := mathskills.SafeRound(mathskills.CalcMedian(data))
	if err != nil {
		fmt.Printf("Error calculating median: %v\n", err)
		return
	}

	// calculate variance and keep it as float for accurate calculations
	variance := mathskills.CalcVariance(data, average)

	// claculate median and round it safely
	standarddeviation, err := mathskills.SafeRound(mathskills.CalcStdDev(variance))
	if err != nil {
		fmt.Printf("Error calculating standard deviation: %v\n", err)
		return
	}

	// round the average safely
	roundedAverage, err := mathskills.SafeRound(average)
	if err != nil {
		fmt.Printf("Error calculating average: %v\n", err)
		return
	}
	// round the variance safely
	roundedVariance, err := mathskills.SafeRound(variance)
	if err != nil {
		fmt.Printf("Error calculating variance: %v\n", err)
		return
	}

	// Print all values to the nearest integer
	fmt.Printf("Average: %d\n", roundedAverage)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", roundedVariance)
	fmt.Printf("Standard Deviation: %d\n", standarddeviation)
}
