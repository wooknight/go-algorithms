package main

import (
	"fmt"
	"math/rand"
)

// Perform linear search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func linearSearch(slice []int, target int) (index, numTests int) {
	for i, value := range slice {
		numTests++
		if value == target {
			return i, numTests
		}
	}
	return -1, numTests
}

func makeRandomSlice(numItems, max int) []int {
	data := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		data[i] = rand.Intn(max)
	}
	return data
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max, target int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)
	slice := makeRandomSlice(numItems, max)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Target: ")
	fmt.Scanln(&target)
	index, numTests := linearSearch(slice, target)
	fmt.Printf("Index: %d, Tests: %d\n", index, numTests)
}
