//given the pos of mth nnode, deletenext n consequtive nodes
// 1-->2-->3-->4--5 pos =1 , n =2 o/p 1-->2-->5
package deleteNnodes

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

func (l *llist) print() {
	temp := l.first
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}

func (l *llist) deleteNnodes(x, y int) error {

	prev := l.first
	for i := 0; i < x; i++ {

		prev = prev.next
	}
	curr := l.first
	for i := 0; i < x+y; i++ {
		curr = curr.next
	}
	prev.next = curr.next

	return nil
}

func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.print()
	err := l.deleteNnodes(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	l.print()
}
