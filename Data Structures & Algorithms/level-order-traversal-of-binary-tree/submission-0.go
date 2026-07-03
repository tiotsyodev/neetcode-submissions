/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, lvl int)
	dfs = func(root *TreeNode, lvl int) {
		if root == nil {
			return
		}

		if lvl == len(res) {
            res = append(res, []int{})
        }
		res[lvl] = append(res[lvl], root.Val)

		dfs(root.Left, lvl + 1)
		dfs(root.Right, lvl + 1)
	}

	dfs(root, 0)

	return res  
    
}
