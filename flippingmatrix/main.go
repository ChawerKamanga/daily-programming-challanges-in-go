package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	n := len(matrix) / 2
	var maxSum int32 = 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Checking all four corresponding positions in the 2n x 2n matrix
			val1 := matrix[i][j]
			val2 := matrix[i][2*n-j-1]
			val3 := matrix[2*n-i-1][j]
			val4 := matrix[2*n-i-1][2*n-j-1]

			// Find the maximum of these four values
			maxVal := max(max(int32(val1), val2), max(val3, val4))

			// Add the maximum value to the sum
			maxSum += maxVal
		}
	}

	return int32(maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var matrix [][]int32
		for i := 0; i < 2*int(n); i++ {
			matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var matrixRow []int32
			for _, matrixRowItem := range matrixRowTemp {
				matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
				checkError(err)
				matrixItem := int32(matrixItemTemp)
				matrixRow = append(matrixRow, matrixItem)
			}

			if len(matrixRow) != 2*int(n) {
				panic("Bad input")
			}

			matrix = append(matrix, matrixRow)
		}

		result := flippingMatrix(matrix)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
