package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.

// traverse the tree, updating the highest visited value in the subtree
// with each visit we check if root.Val is highest than the current highestVal
// if it is, increment goodNodes
func dfs(root *TreeNode, highestVal, goodNodes *int) {
	if root == nil {
		return
	}

	if root.Val >= *highestVal {
		*goodNodes++
		*highestVal = root.Val
	}

	dfs(root.Left, highestVal, goodNodes)
	dfs(root.Right, highestVal, goodNodes)
}

func goodNodes(root *TreeNode) int {
	goodNodes := 0
	highestVal := -99999
	dfs(root, &highestVal, &goodNodes)
	return goodNodes
}
