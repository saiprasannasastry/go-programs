package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type llist struct {
	first  *Node
	last   *Node
	length int
}

func (l *llist) Append(n *Node) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return
	}
	l.last.next = n
	l.last = n
	l.length++
	return
}

func (l *llist) detectloop() bool {
	var s []int
	temp := l.first
	for {
		for _, val := range s {
			if val == temp.value {
				return true
			}
		}
		s = append(s, temp.value)
		temp = temp.next
	}
	return false
	// slowP := l.first
	// fastP := slowP
	// for slowP!=nil && fastP!=nil && fastP.next!=nil {
	// 	slowP = slowP.next
	// 	fastP = fastP.next.next
	// 	if slowP == fastP {
	// 		return true
	// 	}
	// }
	// return false
}
func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.first.next.next.next.next = l.first
	fmt.Println(l.detectloop())
}
