package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, currDepth, minDepth *int) {
	if root == nil {
		return
	}
	*currDepth++
	isLeaf := root.Left == nil && root.Right == nil
	if isLeaf {
		if *minDepth == 0 {
			// set the depth to the first leaf visited
			*minDepth = *currDepth
		} else if *minDepth > *currDepth {
			*minDepth = *currDepth
		}
	}
	dfs(root.Left, currDepth, minDepth)
	dfs(root.Right, currDepth, minDepth)
	*currDepth--
}

func minDepth(root *TreeNode) int {
	currDepth := 0
	minDepth := 0
	dfs(root, &currDepth, &minDepth)
	return minDepth
}
