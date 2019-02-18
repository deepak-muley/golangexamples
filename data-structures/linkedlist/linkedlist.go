package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) PrintLinkedList() {
	for p := ll.head; p != nil; p = p.next {
		fmt.Println(p)
	}
}

func (ll *LinkedList) AddNodeToLinkedList(value int) *Node {
	node := new(Node)
	node.value = value

	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	return node
}

func (ll *LinkedList) InsertAfter(n *Node, value int) *Node {
	newNode := &Node{value: value}
	next := n.next
	n.next = newNode
	newNode.next = next

	return newNode
}

func (ll *LinkedList) ReverseLinkedList() {
	// newHead := ll.head
	// newTail := newHead
	// tmpNode := newHead.next
	//newHead.next =

}
