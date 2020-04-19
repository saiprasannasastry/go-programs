package main

import "fmt"

func main() {
	a := []int{2, 6, 7, 9, 1, 2, 3}
	fmt.Println(minElement(a))
}

func minElement(arr []int) int {
	count := 0

	for range arr {
		count++
	}
	return count
}
