package stringdecompress

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := [5]string{"aabcccccaaa", "aabcdeeqa", "abcd", "a1b2c", "13a12b"}
	for k := range str {
		//fmt.Println(str[k])
		fmt.Println(stringDecomp(str[k]))
	}
}

func stringDecomp(val string) string {
	act_string := ""
	index := 0
	for index < len(val) {
		num := ""
		char := ""
		numb := 1
		for strings.ContainsAny("0123456789", string(val[index])) {
			num += string(val[index])
			index += 1
		}
		if num != "" {
			numb, _ = strconv.Atoi(num)
		}
		char += string(val[index])
		index += 1
		act_string += strings.Repeat(char, numb)

	}
	return act_string
}
