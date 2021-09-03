class TreeNode:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

class BinarySearchTree:
    def __init__(self):
        self.root = None

    def insert(self, data):
        if not self.root:
            self.root = TreeNode(data)
        else:
            self.insertNode(data, self.root)

    def insertNode(self, data, node):
        if data < node.data:
            if node.leftChild:
                self.insertNode(data, node.leftChild)
            else:
                node.leftChild = TreeNode(data)
        else:
            if node.rightChild:
                self.insertNode(data, node.rightChild)
            else:
                node.rightChild = TreeNode(data)

    def getMinValue(self):
        if self.root:
            return self.getMin(self.root)

    def getMin(self, node):
        if node.leftChild:
            return self.getMin(node.leftChild)

        return node.data

    def getMaxValue(self):
        if self.root:
            return self.getMax(self.root)

    def getMax(self, node):
        if node.rightChild:
            return self.getMax(node.rightChild)

        return node.data

    def traverse(self):
        if self.root:
            self.traverseInOrder(self.root)

    def traverseInOrder(self, node):
        if node.leftChild:
            self.traverseInOrder(node.leftChild)

        print("%s ", node.data)

        if node.rightChild:
            self.traverseInOrder(node.rightChild)

    def traversePreOrder(self, node):
        print("%s ", node.data)

        if node.leftChild:
            self.traversePreOrder(node.leftChild)

        if node.rightChild:
            self.traversePreOrder(node.rightChild)

    def traversePostOrder(self, node):
        if node.leftChild:
            self.traversePostOrder(node.leftChild)

        if node.rightChild:
            self.traversePostOrder(node.rightChild)

        print("%s ", node.data)

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
tn.left.left = TreeNode(1)
tn.left.right = TreeNode(3)

print("Preorder order traversal of binary tree:")
traversePreOrder(tn)
print("In order traversal of binary tree:")
traverseInOrder(tn)
print("Postorder order traversal of binary tree:")
traversePostOrder(tn)
print("Level order traversal of binary tree:")
traverseLevelOrder(tn)
