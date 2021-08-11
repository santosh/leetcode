# Binary Search Tree

## Differences from Binary Tree

- The value of left subtree must be lesser than the parent node, and value of the right subtree must be greater than their parent.
- BST are always sorted thus are faster for retrieval.

**Types**: AVL Trees, Red-Black Trees.

## Insertion on BST

We start at the root node. If the data we want to insert is greater than the root node we go to the right, if it is smaller, we go to the left. And so on...

## Traversal on Tree

Tree travarsal refers to the process of visiting (checking and/or updating) each node in a tree data structure, exactly once. Such traversals are classified by the order in which nodes are visited.

Tree traveral can be broadly classified into two categories based on order in which the nodes are visited:

- Breadth-First Search (BFS) Algorithm: It starts from root node and visits every node in current depth before moving to the next depth in the tree.

- Depth-First Search (DFS) Algorithm: It starts with root node and visits all nodes in one branch before backtracking. There are three sub-types under this.

Subtypes of DFS:

- Preorder Traversal - Visits the current node before visiting any nodes inside left or right subtrees.

- Inorder Traversal - Visits the current node after visiting all nodes inside left subtree but before visiting any nde within the right subtree.

- Postorder Traversal - Visit the current node after vising all the nodes in the left and right subtrees.
