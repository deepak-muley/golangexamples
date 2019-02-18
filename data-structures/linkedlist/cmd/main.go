package main

import (
	"fmt"

	"github.com/deepak-muley/golangexamples/data-structures/linkedlist"
)

func main() {
	ll := new(linkedlist.LinkedList)
	ll.AddNodeToLinkedList(10)
	n := ll.AddNodeToLinkedList(20)
	ll.AddNodeToLinkedList(30)
	ll.AddNodeToLinkedList(40)
	ll.PrintLinkedList()
	fmt.Println()

	ll.InsertAfter(n, 25)
	ll.PrintLinkedList()
	fmt.Println()

	ll.ReverseLinkedList()
	ll.PrintLinkedList()
	fmt.Println()
}
