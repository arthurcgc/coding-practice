package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameVal(node1, node2 *TreeNode) bool {
	if node1 == nil {
		if node2 != nil {
			return false
		}
	}
	if node2 == nil {
		return false
	}
	if node1.Val == node2.Val {
		return true
	}
	return false
}

func dfs(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	fmt.Printf("%d\n", root.Val)

	dfs(root.Left)
	dfs(root.Right)
	return false
}

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	for dfs(t1) == true {
		fmt.Printf("halt\n")
	}
}
