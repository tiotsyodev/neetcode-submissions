/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := []int{}

	var dfs func(node *TreeNode, lvl int)
	dfs = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}

		if lvl == len(res) {
			res = append(res, node.Val)
		}

		dfs(node.Right, lvl + 1)
		dfs(node.Left, lvl + 1)
	}

	dfs(root, 0)

	return res
}
