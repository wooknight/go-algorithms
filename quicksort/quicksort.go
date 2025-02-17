package main

import "fmt"

func main() {
	data := []int{1, 12, 3, 4, 5, 6, 7, 8, 9, 10, 11, 2, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	fmt.Println(data)
	quicksort(data, 0, len(data)-1)
	fmt.Println(data)
}

func quicksort(data []int, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		quicksort(data, low, pivot-1)
		quicksort(data, pivot+1, high)
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
