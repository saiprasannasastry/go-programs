package main

import "fmt"

func maxArray(arr []int) []int {

	var s []int
	for i := 0; i < len(arr); i++ {
		initialVal := arr[i]
		sum := initialVal
		for j := i + 1; j < len(arr); j++ {
			sum = sum + arr[j]
			if sum == 0 {
				return arr[i : j+1]
			}

		}
	}
	return s
}

func main() {
	a := []int{9, -7, -5, 0, -1, 5, 8, 11}
	fmt.Println(maxArray(a))
	a = []int{6, -6, 8, 4, -12, 9, 8, -8}
	fmt.Println(maxArray(a))
	//fmt.Println(a)
}
