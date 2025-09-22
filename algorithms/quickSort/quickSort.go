package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	sortedLeft := quicksort(left)
	sortedRight := quicksort(right)

	return append(append(sortedLeft, pivot), sortedRight...)
}

func main() {
	arr := []int{19, 5, 3, 12, 7, 8, 1, 17, 11, 6, 4, 14, 2, 16, 9, 15, 10, 13, 18}
	fmt.Println("Original:", arr)
	sorted := quicksort(arr)
	fmt.Println("Ordenado:", sorted)
}
