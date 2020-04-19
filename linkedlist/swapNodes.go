//swap node x ,y 1-->2-->3-->4--5 , swap 3 and 4 .
//OP 1-->2-->4-->3--5
package swapNodes

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

func (l *llist) swapNodes(x, y int) error {
	prevX := l.last.next
	currX := l.first
	for currX != nil && currX.value != x {
		prevX = currX
		currX = currX.next
	}

	prevY := l.last.next
	currY := l.first
	for currY != nil && currY.value != y {
		prevY = currY
		currY = currY.next
	}
	if currX == nil || currY == nil {
		return fmt.Errorf("nothing to swap")
	}
	if prevX != nil {
		prevX.next = currY
	} else {
		l.first = currY
	}

	if prevY != nil {
		prevY.next = currX
	} else {
		l.first = currX
	}
	temp := currX.next
	currX.next = currY.next
	currY.next = temp

	return nil
}

func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.print()
	err := l.swapNodes(3, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	l.print()
}
