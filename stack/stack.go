package main

import "fmt"

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(x string) {
	*s = append(*s, x)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		pos := len(*s) - 1
		ele := (*s)[pos]
		*s = (*s)[:pos]
		return ele, true
	}

}

func main() {
	var s stack

	fmt.Println(s.isEmpty())
	s.push("hi")
	s.push("hi1")
	s.push("hi2")
	s.push("hi3")
	fmt.Println(s.pop())
}
