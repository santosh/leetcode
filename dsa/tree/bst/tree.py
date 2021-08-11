class TreeNode:
    def __init__(self, data) -> None:
        self.data = data
        self.left = None
        self.right = None

    def insert(self, node):
        pass

def traversePreOrder(root):
    if root:
        print(root.data)
        traversePreOrder(root.left)
        traversePreOrder(root.right)

def traverseInOrder(root):
    if root:
        traverseInOrder(root.left)
        print(root.data)
        traverseInOrder(root.right)

def traversePostOrder(root):
    if root:
        traversePostOrder(root.left)
        traversePostOrder(root.right)
        print(root.data)

def traverseLevelOrder(root):
    if root is None:
        return

    queue = []

    queue.append(root)

    while len(queue) > 0:
        print(queue[0].data)
        node = queue.pop(0)

        if node.left is not None:
            queue.append(node.left)

        if node.right is not None:
            queue.append(node.right)


tn = TreeNode(10)
tn.left = TreeNode(5)
tn.right = TreeNode(15)
tn.left.left = TreeNode(3)
tn.left.right = TreeNode(1)

print("Preorder order traversal of binary tree:")
traversePreOrder(tn)
print("In order traversal of binary tree:")
traverseInOrder(tn)
print("Postorder order traversal of binary tree:")
traversePostOrder(tn)
print("Level order traversal of binary tree:")
traverseLevelOrder(tn)
