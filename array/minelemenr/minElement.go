package main

import "fmt"

func main() {
	a := []int{2, 6, 7, 9, 1, 2, 3}
	fmt.Println(minElement(a))
}

func minElement(arr []int) int {
	min := arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
