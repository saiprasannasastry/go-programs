package stringcompress

import (
	"fmt"
	"strconv"
)

func main() {
	str := [4]string{"aabcccccaaa", "aabcdeeqa", "abcd", "acdeeewqa"}
	for k := range str {
		//fmt.Println(str[k])
		fmt.Println(stringcomp(str[k]))
	}
}

func stringcomp(val string) string {
	comp_str := ""
	count := 1
	for k := 0; k < len(val)-1; k++ {
		if string(val[k]) == string(val[k+1]) {
			count += 1
		} else {
			comp_str += string(val[k]) + strconv.Itoa(count)
			count = 1
		}
	}
	comp_str += string(val[len(val)-1]) + strconv.Itoa(count)

	if len(comp_str) >= len(val) {
		return val
	}
	return comp_str

}
