/*Given an array A of size N, your task is to do some operations,
 i.e., search an element, insert an element, and delete an element
  by completing the functions.
 Also, all functions should return a boolean value

 Input Format:
The first line of input consists of T, the number of test cases. T testcases follow. Each testcase contains 3 lines of input. The first line of every test case consists of an integer N, denotes the number of elements in an array. The second line of every test cases consists of N spaced integers Ai. The third line of every test case consists of four integers X, Y, Yi and Z, denoting the elements for the search operation, insert operation, insert operation element index and delete operation respectively.

Output Format:
For each testcase, return 1 for each operation if done successfully else 0.

Your Task:
Since this is a function problem, you only need to complete the provided functions.

Constraints:
1 <= T <= 100
1 <= N <= 100
1 <= Ai <= 100

Example:
Input:
1
5
2 4 1 0 6
1 2 2 0
Output:
1 1 1
*/

package task1

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}

	fmt.Println(search(a, 5))
	fmt.Println(insert_ele(a, 7, 6))
	fmt.Println(delete_ele(a, 4))
}

func search(a []int, value int) bool {

	for _, a := range a {
		if a == value {
			return true
		}
	}

	return false
}

func insert_ele(arr []int, value int, pos int) []int {
	if pos >= len(arr) {
		pos = len(arr)
	}
	out := make([]int, len(arr)+1)
	copy(out[:pos], arr[:pos])
	out[pos] = value
	copy(out[pos+1:], arr[pos:])
	return out
}

func delete_ele(arr []int, value int) []int {
	new_arr := []int{}
	for _, a := range arr {
		if a != value {
			new_arr = append(new_arr, a)
		}
	}

	return new_arr
}
