package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Convert the input string to lowercase for case insensitivity
	s = strings.ToLower(s)

	// Create a map to hold the unique letters of the alphabet
	alphabetMap := make(map[rune]bool)
	for r := 'a'; r <= 'z'; r++ {
		alphabetMap[r] = true
	}

	// Iterate over each character in the input string
	for _, char := range s {
		// Check if the character is a letter and remove it from the map
		if _, ok := alphabetMap[char]; ok {
			delete(alphabetMap, char)
		}
	}

	// If the map is empty, all letters were present, return "pangram"
	if len(alphabetMap) == 0 {
		return "pangram"
	}

	// If the map is not empty, some letters were missing, return "not pangram"
	return "not pangram"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

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
