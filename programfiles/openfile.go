package mathskills

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(filePath string) ([]float64, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file")
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var data []float64
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert string to float: %v", err)
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
