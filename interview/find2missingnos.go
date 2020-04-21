//given an array find the 2 missing numbers
package findMissingNos

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 4, 6}
	fmt.Println(findMissingNos(a))
}

func findMissingNos(a []int) []int {
	sort.Ints(a)
	boolArr := make([]bool, a[len(a)-1])
	newArr := []int{}
	for _, v := range a {
		boolArr[v-1] = true
	}

	for k, v := range boolArr {
		if v == false {
			newArr = append(newArr, k)
		}
	}

	return newArr
}
