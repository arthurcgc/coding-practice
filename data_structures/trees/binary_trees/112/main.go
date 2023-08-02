package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

func dfs(root *TreeNode, targetSum int, currentCost *int, hasMatchingCost *bool) {
	if root == nil || *hasMatchingCost {
		return
	}

	*currentCost += root.Val
	if root.Left == root.Right && root.Left == nil {
		// means we're on a leaf, so we should check the targetSum
		if targetSum == *currentCost {
			*hasMatchingCost = true
			return
		}
	}

	dfs(root.Left, targetSum, currentCost, hasMatchingCost)
	dfs(root.Right, targetSum, currentCost, hasMatchingCost)
	// if we're here, means we're going up, so we have to decrement root.Val
	*currentCost -= root.Val
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	currentCost := 0
	hasMatchingCost := false
	dfs(root, targetSum, &currentCost, &hasMatchingCost)
	return hasMatchingCost
}

func main() {

}
