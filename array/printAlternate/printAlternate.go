package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 8}
	fmt.Println(alternate(a))
}

func alternate(arr []int) []int {
	count := 0
	newArr := make([]int, (len(arr)+1)/2)
	for _, val := range arr {
		if count%2 == 0 {
			newArr[count/2] = val
		}
		count++
	}
	return newArr
}
