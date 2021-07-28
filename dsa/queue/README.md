# Queue

- It is an abstract data type.
- Basic operations: enqueue() & dequeue(), peek()
- FIFO structure: first in first out
- It can be implemented with dynamic arrays as well as with linked lists
- Important when implementing BFS algorithms for graphs

Enqueue operation: We just simply add the new items to the end of the queue. It is an O(1) operation.

Dequeue operation: we just simply remove the item starting at the beginning of the queue // FIFO structure

## Applications

- When a resource is shared with several consumers (threads): we store them in a queue. For example: CPU scheduling
- When data is transferred asynchronously (data not necessarily received at same rate as sent) between two processes. For example: IO buffers
- Operational research applications or stochastic models relies heavily on queues!!!
