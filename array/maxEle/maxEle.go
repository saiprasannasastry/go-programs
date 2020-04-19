package main

import "fmt"

func main() {
	a := []int{2, 6, 7, 9, 1, 2, 3}
	fmt.Println(maxEle(a))
}

func maxEle(arr []int) int {
	max := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}
