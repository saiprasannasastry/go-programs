// Given an array arr of size N and an integer K, 
// the task is to find the maximum for each and every contiguous subarray of size K.
// Input: arr[] = {1, 2, 3, 1, 4, 5, 2, 3, 6}, K = 3
// Output: 3 3 4 5 5 5 6

package maxSlidingWindow

import "fmt"

func main(){

	arr := []int{1,4,3,1,4,5,2,3,6}
	k :=3
	fmt.Println(maxSlidingWindow(arr,k))
}

func maxSlidingWindow(arr []int, k int) []int{
	newArr := make([]int,0)
	if k==1{
		return arr
	}

	initialPointer :=0
	max :=arr[k-1]
	q := k-1
	t := initialPointer

	for q <= len(arr)-1{
		if arr[initialPointer] > max{
			max = arr[initialPointer]
		}
		initialPointer +=1

		if q == initialPointer && initialPointer != len(arr) {
			newArr = append(newArr, max)
			q += 1
			initialPointer = t + 1
			t=p 

			if q < len(arr) {
				max = arr[q]
			}
		}
	}
	return newArr
}