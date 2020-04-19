//bubble sdort as well
package main

import "fmt"

func main() {
	arr := []int{2, 7, 6, 1, 5}
	n := 2
	fmt.Println(kthsmallest(arr, n))
}

func kthsmallest(a []int, n int) int {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a[n-1]
}
