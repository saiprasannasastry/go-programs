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

func (l *llist) Prepend(n *Node) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return

	}
	first_element := l.first
	l.first = n
	l.first.next = first_element
	l.length++
	return
}
func (l *llist) Lookup() []int {
	if l.length == 0 {
		return nil
	}
	var nodes []int
	node := l.first
	for node != nil {
		nodes = append(nodes, node.value)
		node = node.next
	}
	return nodes
}

func (l *llist) Delete(n int) {
	if l.length == 0 {
		return
	}
	if l.first.value == n {
		l.first = l.first.next
		l.length--
		return
	}
	var prevNode *Node
	node := l.first
	for node.value != n {
		if node == nil {
			return
		}
		prevNode = node
		node = node.next
	}
	prevNode.next = node.next
	l.length--
}

func main() {
	l := &llist{}

	secondNode := &Node{
		value: 2,
	}
	l.Append(secondNode)

	thirdNode := &Node{
		value: 3,
	}
	l.Append(thirdNode)

	firstNode := &Node{
		value: 1,
	}
	l.Prepend(firstNode)

	fmt.Println(*l.first)
	fmt.Println(l.first.next)
	fmt.Println(l.last)
	fmt.Println(l.length)
	fmt.Println(l.Lookup())

	l.Delete(1)
	fmt.Println(l.Lookup())

	l.Delete(3)
	fmt.Println(l.Lookup())
}
