package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, greatest int, goods *int){
    if root == nil {
        return
    }

    if root.Val >= greatest {
        *goods++
        if root.Left != nil || root.Right != nil {
            greatest = root.Val
        }
    }
    dfs(root.Left, greatest, goods)
    dfs(root.Right, greatest, goods)
}

func goodNodes(root *TreeNode) int {
    goods := 0
    dfs(root, -999999, &goods)
    return goods
}
