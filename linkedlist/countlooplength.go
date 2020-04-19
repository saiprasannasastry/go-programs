//1-->2-->3-->4-->5-->2-->3-->4--5>--2> .......

package countLoopLenght

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

func (l *llist) countLoopLenght() int {
	//var s []int
	//temp := l.first
	count := 0
	flag := 0
	slowP := l.first
	fastP := slowP
	for slowP != nil && fastP != nil && fastP.next != nil {
		if slowP.value == fastP.value && flag != 0 {
			count++
			slowP = slowP.next
			for slowP.value != fastP.value {
				slowP = slowP.next
				count += 1
			}
			return count
		}
		slowP = slowP.next
		fastP = fastP.next.next
		flag = 1
	}
	return 0
}
func main() {
	l := &llist{}

	for i := 1; i <= 5; i++ {

		l.Append(&Node{value: i})
	}
	l.first.next.next.next.next = l.first
	fmt.Println(l.countLoopLenght())
}
