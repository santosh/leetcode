package main

import "fmt"

type Node struct {
    val interface{}
    next *Node
}

type LinkedList struct {
    head    *Node
    len     int
}

// append is an alias for insertAtEnd
func (ll *LinkedList) append(node *Node) {
	ll.insertAtEnd(node)
}

func main() {
	myList := LinkedList{}

	myList.head = &Node{1, nil}
	myList.len = 1

	myList.traverse()
}

// insertAtEnd inserts a Node at the end of the list.
func (ll *LinkedList) insertAtEnd(node *Node) {

}

// insertAtIndex inserts a Node at given index. 
func (ll *LinkedList) insertAtIndex() {

}

// traverse methods goes through each node from first to last 
// and prints each node's value
func (ll *LinkedList) traverse() {
	fmt.Printf("Traversing: ")
	fmt.Println(ll.head.val)
}

func (ll *LinkedList) search(val interface{}) {
	// traverse and print each node till end
	ll.traverse()
}

// type LinkedList interface {
// 	   get()
// 	   access()
//     traverse()
//     insertAtIndex()
//     insertAtStart()
//     insertAtEnd()
//     deleteFromIndex()
//     deleteFromStart()
//     deleteFromEnd()
//	   search(val interface{})
// }
