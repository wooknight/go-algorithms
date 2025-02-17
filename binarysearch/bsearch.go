package main

import (
	"fmt"
	"math/rand"
)

// Perform binary search.
// Return the target's location in the slice and the number of tests.
// If the item is not found, return -1 and the number tests.
func binarySearch(slice []int, target int) (index, numTests int) {
	low := 0
	high := len(slice) - 1
	for low <= high {
		numTests++
		mid := (low + high) / 2
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(slice) || slice[low] != target {
		return -1, numTests
	}
	return low, numTests
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
	InnerQuicksort(slice, 0, len(slice)-1)
	fmt.Printf("Sorted: %v\n", slice)
	fmt.Printf("Target: ")
	fmt.Scanln(&target)
	index, numTests := binarySearch(slice, target)
	fmt.Printf("Index: %d, Tests: %d\n", index, numTests)
}

func InnerQuicksort(data []int, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		InnerQuicksort(data, low, pivot-1)
		InnerQuicksort(data, pivot+1, high)
	}
}

func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low - 1
	for j := low; j < high; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}
