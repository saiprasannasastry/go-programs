package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	N := len(arr)

	var n int
	fmt.Scanln(&n)
	fmt.Println(rotArr(arr, N, n))
}

func rotArr(arr []int, size int, rotateby int) []int {
	for i := 0; i < rotateby; i++ {
		leftrotate(arr, size)
	}
	return arr
}

func leftrotate(arr []int, size int) []int {
	tmp := arr[0]

	for i := 0; i < size-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[size-1] = tmp
	return arr
}
