//1-->2-->4-->3--5
// o/p 5-->4-->3-->2-->1
package reverse

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
func (l *llist) reverse() {
	curr := l.first
	prev := l.last.next
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.first = prev

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
	(l.reverse())
	fmt.Println()
	l.print()
}
