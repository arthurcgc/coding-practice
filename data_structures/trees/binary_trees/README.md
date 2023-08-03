# Introduction
There are two main types of tree traversals. The first is called depth-first search (DFS). For binary trees specifically, there are 3 ways to perform DFS - preorder, inorder, and postorder (don't worry though, the type you choose rarely matters). The other main type of traversal is called breadth-first search (BFS). Let's start by looking at DFS.

## DFS
> Recall that the depth of a node is its distance from the root. 

In a DFS, we prioritize depth by traversing as far down the tree as possible in one direction (until reaching a leaf node) before considering the other direction. For example, let's say we choose left as our priority direction. We move exclusively with node.left until the left subtree has been fully explored. Then, we explore the right subtree.

Once it fully explores the branch, it backtracks until it finds another unexplored branch.

### Pseudo/Python Code
```python
def dfs(node):
    if node == None:
        return
    
    dfs(node.left)
    dfs(node.right)
    return
```

* Preorder DFS
    In preorder traversal, logic is done on the current node before moving to the children.
    ```python
        def dfs(node):
        if node == None:
            return
        
        # do some logic

        dfs(node.left)
        dfs(node.right)
        return
    ```
* Inorder DFS
    For inorder traversal, we first recursively call the left child, then perform logic on the current node, then recursively call the right child. This means no logic will be done until we reach a node without a left child since calling on the left child takes priority over performing logic.
    ```python
        def dfs(node):
        if node == None:
            return

        dfs(node.left)
        # do some logic
        dfs(node.right)
        return
    ```
* Postorder DFS
    In postorder traversal, we recursively call on the children first and then perform logic on the current node. This means no logic will be done until we reach a leaf node since calling on the children takes priority over performing logic. In a postorder traversal, the root is the last node where logic is done.
    ```python
        def dfs(node):
        if node == None:
            return

        dfs(node.left)
        dfs(node.right)
        # do some logic
        return
    ```
## BFS
    In DFS, we prioritized depth. In breadth-first search (BFS), we prioritize breadth. Recall that a node's depth is its distance from the root. In DFS, we always tried to go down as far as we could, increasing the depth of the current node until we reached a leaf.
    In BFS, we traverse all nodes at a given depth before moving on to the next depth.

    While DFS was implemented using a stack (recursion uses a stack under the hood), BFS is implemented iteratively with a queue. You can implement BFS with recursion, but it wouldn't make sense as it's a lot more difficult without any benefit. As such, we will look only at iterative implementations.

    
