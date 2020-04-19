package main

import "fmt"

func main() {
	sum := 16
	arr := []int{1, 11, 5, 12, 4}
	fmt.Println(findPairinArr(arr, sum))
}

func findPairinArr(a []int, sum int) map[int]int {
	m1 := make(map[int]int)

	for k, _ := range a {
		j := k + 1
		for j, _ = range a {
			if a[k]+a[j] == sum {
				m1[a[k]] = a[j]

			}
		}
	}
	return m1
}
