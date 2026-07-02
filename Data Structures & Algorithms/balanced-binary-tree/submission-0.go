/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	res := true

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil || !res {
			return 0 
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		if int(math.Abs(float64(l - r))) > 1 {
			res = false
			return 0
		}

		return 1 + max(l, r)
	}

	dfs(root)

	return res
    
}
