//given a sorted linked list, delete the duplicates

package deleteDuplicate

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
func (l *llist) deleteDuplicate() {
	temp := l.first
	//fmt.Print(temp)
	for temp.next != nil {
		if temp.value == temp.next.value {
			new := temp.next.next
			temp.next = l.last.next
			temp.next = new

		} else {
			temp = temp.next
		}
	}
}
func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.Append(&Node{value: 5})
	l.Append(&Node{value: 6})
	l.Append(&Node{value: 7})
	l.Append(&Node{value: 7})
	l.print()

	l.deleteDuplicate()
	fmt.Println()
	l.print()
}
