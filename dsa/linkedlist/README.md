# Linked List

- Is not pre-initilized. The size of the linked list is not declared first, in contrast to arrays, where we have to fix the size of the array.
- Stores each list item in non contagious memory location. Thus there is no wastage of memory 

There are at least two variables which should be there in a linked list node. 

- The body of the node. 
- Pointer to the next node.

```go
type Node struct {
    val interface{}
    next *Node
}
```

Our linked list is composed of these `Node`s. A linked list can be without any node. Let's see that with an example. On top of that, a linked list should be able to perform these operation:

- traversal
- insertion
- deletion

We'll create an interface for our Linked List. 

```go
type LinkedList struct {
    head    *Node
    len     int
}
```
