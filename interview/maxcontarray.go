//given an array find the max contiguous subarray
package maxContarray

import (
	"fmt"
	"math"
)

func maxContarray(arr []float64, len int) float64 {
	max_so_far := arr[0]
	curr_max := arr[0]
	for i := 1; i < len; i++ {
		curr_max = math.Max(arr[i], curr_max+arr[i])
		max_so_far = math.Max(max_so_far, curr_max)
	}
	return max_so_far
}

func main() {
	a := []float64{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(maxContarray(a, len(a)))
}
