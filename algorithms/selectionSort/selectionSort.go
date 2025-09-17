package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{2, 1, 7, 3, 23, 9, 8, 6, 22, 32, 24}))
}

func searchLowest(arr []int) int {
	menor := arr[0]
	menorIndice := 0
	for i := range arr {
		if arr[i] < menor {
			menor = arr[i]
			menorIndice = i
		}
	}
	return menorIndice
}

func selectionSort(arr []int) []int {
	novoArr := []int{}

	for range arr {
		menor := searchLowest(arr)
		novoArr = append(novoArr, arr[menor])
		arr = append(arr[:menor], arr[menor+1:]...)
	}
	return novoArr
}
