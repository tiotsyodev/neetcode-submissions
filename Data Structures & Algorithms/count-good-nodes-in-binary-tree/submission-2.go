/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode, b int)
	dfs = func(node *TreeNode, b int) {
		if node == nil {
			return 
		} 

		if node.Val >= b {
			count++
			b = node.Val
		} 

		dfs(node.Right, b)
		dfs(node.Left, b)
	}

	dfs(root, root.Val)

	return count
    
}
