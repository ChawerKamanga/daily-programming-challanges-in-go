package main

import "fmt"

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	seenSocks := make(map[int32]int32) // Create a map to store seen colors and their first occurrence index

	pairs := int32(0) // Initialize number of pairs

	for _, color := range ar {
		if _, ok := seenSocks[color]; ok { // Check if color exists in the map
			pairs++ // Increment pairs and remove first occurrence
			delete(seenSocks, color)
		} else {
			seenSocks[color] = color // Add color and its index to the map
		}
	}
	fmt.Println(seenSocks)

	return pairs
}

func main() {
	// Test case 1
	n1 := int32(7)
	ar1 := []int32{1, 2, 1, 2, 1, 3, 2}
	expectedPairs1 := int32(2)

	pairs1 := sockMerchant(n1, ar1)

	if pairs1 != expectedPairs1 {
		fmt.Println("Test case 1 failed. Expected:", expectedPairs1, "Got:", pairs1)
	} else {
		fmt.Println("Test case 1 passed.")
	}

	// Add more test cases here if needed...
}
