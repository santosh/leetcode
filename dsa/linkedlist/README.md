# Linked List

- Linked lists are composed of nodes and references / pointing from one node to the other!!
- The last reference is pointing to a NULL.
- Is not pre-initilized. The size of the linked list is not declared first, in contrast to arrays, where we have to fix the size of the array.
- Stores each list item in non contagious memory location. Thus there is no wastage of memory 
- They can be used to implement several other common data types: stacks, queues

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

Operations on Linked List:

1. add

Takes `O(1)` when inserting in front. Will take upto `O(n)` when inserting in between. Inserting at end is always `O(n)`.

2. remove

Takes `O(1)` when removing from front. Will take upto `O(n)` when removing from between. Removing from end is always `O(n)`.

3. search/traverse

Takes `O(n)` when using linear search.


