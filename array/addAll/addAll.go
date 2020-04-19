package main

import "fmt"

func main() {
	a := []int{2, 6, 7, 9, 1, 2, 3}
	fmt.Println(addAll(a))
}

func addAll(arr []int) int {
	val := 0

	for _, value := range arr {
		val += value
	}
	return val
}
