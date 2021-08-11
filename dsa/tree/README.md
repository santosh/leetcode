# Tree

Tree is a hierarchical data structure which stores the information naturally in the form of hierarchy unlike linear data structures like, Linked List, Stack etc. A tree contains nodes (data) and connections (edges) which should not form a cycle.

Note: In a tree, there must be only a single path from the root node te any other nodes in the tree. If there are several ways to get to a given node, then it is not a tree.

**Node** - A node is a structure which may contain a value or condition, or represent a separate data structure.

**Root** - The top node in a tree, the prime ancestor.

**Child** - A node directly connected to another node when moving away from the root, an immediate descendant.

**Parent** - The converse notion of a child, an immediate ancestor. 

**Leaf** - A node with no children.

**Internal node** - A node with a least one child.

**Edge** - The connection between one node and another.

**Depth** - This distance between a node and the root.

**Level** - The number of edges between a node and the root + 1

**Height** - The number of edges on the longest path between a node and a descendant leaf.

**Breadth** - The number of leaves.

**Sub Tree** - A tree T is a tree consisting of a node in T and all of its descendants in T.

**Binary Tree** - A tree data structure in which each node has at most two children, which are referred to as the left child and the right child.

**Binary Search Tree** - It is a special type of binary tree which has the following properties.

- The left subtree of a node contains only nodes with keys lesser than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- The left and right subtree each must also be a binary search tree.

## Traversal on Tree

See [`bst/README.md`](bst/README.md).
