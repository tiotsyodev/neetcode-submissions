/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
        return false
    }

	sl := head
	fs := head.Next

	for sl != fs {
		if fs == nil || fs.Next == nil {
			return false
		}

		sl = sl.Next
		fs = fs.Next.Next
	}

	return true
}
