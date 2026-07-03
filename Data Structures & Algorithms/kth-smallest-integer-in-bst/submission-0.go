/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := []int{}
	inorder(root, &res)
	return res[k - 1]
    
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
