package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
func dfs(root *TreeNode, currDepth, maxDepth *int) {
	if root == nil {
		return
	}

	*currDepth++
	if *currDepth > *maxDepth {
		*maxDepth = *currDepth
	}

	dfs(root.Left, currDepth, maxDepth)
	dfs(root.Right, currDepth, maxDepth)
	// visited all the leaves, that means we're going up a level
	*currDepth--
}

func maxDepth(root *TreeNode) int {
	currD := 0
	maxD := 0
	dfs(root, &currD, &maxD)
	return maxD
}

func main() {

}
