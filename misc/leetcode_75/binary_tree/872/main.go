package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, s *string) {
    if root == nil {
        return
    }

    dfs(root.Left, s)
    dfs(root.Right, s)
    if root.Left == nil && root.Right == nil {
        *s += strconv.Itoa(root.Val) + "-"
    }
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var s1 string
    var s2 string
    dfs(root1, &s1)
    dfs(root2, &s2)
    return s1 == s2
}
