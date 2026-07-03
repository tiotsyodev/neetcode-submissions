/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    cur := root
    var prev *int
    
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if prev != nil && cur.Val <= *prev {
            return false
        }
        prev = &cur.Val
        
        cur = cur.Right
    }
    
    return true
}
