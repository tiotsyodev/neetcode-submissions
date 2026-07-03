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
    
	var dfs func(node *TreeNode, lvl int)
	dfs = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}

		if len(res) == lvl {
			res = append(res, []int{})
		}

		res[lvl] = append(res[lvl], node.Val)
		
		dfs(node.Left, lvl + 1)
		dfs(node.Right, lvl + 1)
	}

	dfs(root, 0)
	
	return res
}
