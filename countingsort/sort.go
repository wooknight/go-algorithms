package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	data := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		data[i] = rand.Intn(max)
	}
	return data
}
func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func countingSort(slice []int, max int) []int {
	counts := make([]int, max+1)
	for _, item := range slice {
		counts[item]++
	}
	sorted := make([]int, len(slice))
	index := 0
	for i, count := range counts {
		for j := 0; j < count; j++ {
			sorted[index] = i
			index++
		}
	}
	return sorted
}

func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("*** Not sorted!")
			return
		}
	}
	fmt.Println("Sorted!")
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}
