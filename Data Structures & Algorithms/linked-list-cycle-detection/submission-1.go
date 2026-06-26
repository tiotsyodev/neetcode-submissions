/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
        return false
    }

	mp := map[*ListNode]int{}
    cur := head

	for cur != nil {
		_, ok := mp[cur]
		
		if ok {
			return true
		}

		mp[cur]++
		cur = cur.Next
	}

	return false
}
