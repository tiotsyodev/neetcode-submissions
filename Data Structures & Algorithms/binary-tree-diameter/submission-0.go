/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {

	counter := 0

	var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := dfs(node.Left)
        rightHeight := dfs(node.Right)

		if rightHeight + leftHeight > counter {
			counter = rightHeight + leftHeight
		}

		return 1 + max(leftHeight, rightHeight)

	}

	dfs(root)
	return counter 
    
}
