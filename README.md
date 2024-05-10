# math-skills

## Overview
This Go program calculates essential statistical measures such as Average, Median, Variance, and Standard Deviation from a dataset provided in a text file. Each line of the file represents an individual data point of a statistical population.

### Instructions
The program reads from a file and print the result of each statistic . The program reads the data present in the path passed as argument. The data in the file will be presented as the following example:

189
113
121
114
145
110
...

This data represents a statistical population: each line contains one value.


## Features
- **Reads data from a file**: The program reads numbers from a file where each line contains one numerical data point.
- **Calculates Statistical Measures**:
  - **Average**: Computes the mean of the dataset.
  - **Median**: Determines the middle value of the dataset.
  - **Variance**: Calculates the measure of the spread of numbers in the dataset.
  - **Standard Deviation**: Computes the dispersion of the dataset from the mean.
- **Outputs Rounded Results**: All results are rounded to the nearest integer and printed to the console.

## Prerequisites
You need to have Go installed on your system to run this program. Go can be downloaded from [https://golang.org/dl/](https://golang.org/dl/).

## Setup
To set up the project on your local machine, follow these steps:

```bash
git clone https://learn.zone01kisumu.ke/git/shfana/math-skills
cd math-skills
```
# Usage

To run the program, you will pass the path of the data file as a command-line argument. For example:

```bash

go run main.go data.txt

```
This command will process the numbers in data.txt, compute the required statistical measures, and output them rounded to the nearest integer as shown below:

```bash

Average: <average>
Median: <median>
Variance: <variance>
Standard Deviation: <standard deviation>
```
## Testing

### Overview

Tests are organized to validate each statistical function in the mathskills package. This ensures accuracy and robustness of the computations.
Running Tests

Execute the following command in the terminal to run all tests:

```bash

go test -v ./mathskills/...

```

To run tests for a specific function:

```bash

go test -v ./mathskills/average_test.go
```

Expected Test Output

Successful test execution will display:

```diff

=== RUN   TestCalcAverage
--- PASS: TestCalcAverage (0.00s)
=== RUN   TestCalcMedian
--- PASS: TestCalcMedian (0.00s)
=== RUN   TestCalcVariance
--- PASS: TestCalcVariance (0.00s)
=== RUN   TestCalcStdDev
--- PASS: TestCalcStdDev (0.00s)
PASS
ok      mathskills 0.005s
```
### Learning Outcomes

  - Statistics and Mathematics: Understanding and applying basic statistical calculations.
  - File Handling in Go: Reading and processing data effectively using Go.
  - Robust Go Programming: Implementing and testing reliable Go programs.

### Contributions

Contributions are welcome. Please ensure your code adheres to the existing style and includes tests where applicable.
 
 ## License 
 This project is protected under the MIT Licence. See the license for details

 ## Author
 This project was created by [shfana](https://github.com/Wambita)
