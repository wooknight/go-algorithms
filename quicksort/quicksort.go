package main

import (
	"fmt"
	"math/rand"
)

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
	quicksort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}

func makeRandomSlice(numItems, max int) []int {
	data := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		data[i] = rand.Intn(max)
	}
	return data
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

func quicksort(data []int) {
	InnerQuicksort(data, 0, len(data)-1)
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
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
