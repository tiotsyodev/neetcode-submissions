/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {

	diameter := 0 

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil { 
			return 0
		}

		r := dfs(node.Right)
		l := dfs(node.Left)

		if l + r > diameter {
			diameter = l + r
		}

		return 1 + max(l, r)
	}

	dfs(root)

	return diameter
    
}
