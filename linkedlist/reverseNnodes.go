//given a linked list reverse the N given nodes
//1-->2-->3-->4-->5 n=3 o/p 3-->2--1-->4-->5

package reverseNnodes

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
func (l *llist) reverseNnodes(n int) {
	curr := l.first
	prev := l.last.next
	count := 0
	for curr != nil && count != n {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
		count += 1

	}
	l.first = prev
	tmp := l.first
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = curr

}
func (l *llist) print() {
	temp := l.first
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}

func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.print()
	fmt.Println()
	n := 4
	(l.reverseNnodes(n))
	fmt.Println()
	l.print()
}
