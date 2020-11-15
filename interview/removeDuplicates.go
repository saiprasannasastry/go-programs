//given an array , remove duplicates from the arrau
//sol1 store a Temp array, O(N) Time, O(N )space
//below is an O(1) Space implementation and O(NlogN)
//using go sort package, if not use mergeSortAlgo first

package removeDuplicates

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 3, 2, 1, 4, 6, 5, 5, 7, 88}
	removeDuplicates(arr)
}

func removeDuplicates(arr []int) {
	j := 0
	sort.Ints(arr)
	fmt.Println(arr)
	for k := 0; k < len(arr)-1; k++ {
		if arr[k] != arr[k+1] {
			arr[j] = arr[k]
			j++
		}
	}
	arr[j] = arr[len(arr)-1]

	j++
	arr = arr[:j]
	fmt.Println(arr)
}
