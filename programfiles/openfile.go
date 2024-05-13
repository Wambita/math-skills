package mathskills

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func OpenFile(filePath string) ([]float64, error) {
	// Opening the file
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file")
	}
	defer inputFile.Close()

	// check for empty file
	fileinfo, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("os error reading file")
	}
	if fileinfo.Size() == 0 {
		return nil, fmt.Errorf("empty file")
	}

	// Slice of floats to store float values
	// scanner object to read, convert to float and append to slice
	var data []float64
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if num > float64(math.MaxInt64) {
			return nil, fmt.Errorf("value out of int64 range")
		}
		if num < 0 {
			return nil, fmt.Errorf("population cannot have a negative value")
		}
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		data = append(data, num)

	}
	if len(data) < 2 {
		return nil, fmt.Errorf("population must have more than 1 value")
	}
	return data, nil
}
