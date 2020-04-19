package main

import "fmt"

func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--

		}
		arr[j+1] = key
	}
}

func main() {
	a := []int{13, 12, 11, 5, 6}
	insertionSort(a)
	fmt.Print(a)
}
