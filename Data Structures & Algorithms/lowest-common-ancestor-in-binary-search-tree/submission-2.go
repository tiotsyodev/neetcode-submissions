/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {

		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}

	}

	return nil
}
