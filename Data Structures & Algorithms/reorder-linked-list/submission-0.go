/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	center := slow
	var prev *ListNode

	for center != nil {
		next := center.Next
		center.Next = prev
		prev = center
		center = next
	}

	cur1 := head
	cur2 := prev
	
	for cur1 != nil && cur2 != nil {

		next1 := cur1.Next
		next2 := cur2.Next

		cur1.Next = cur2
		cur2.Next = next1


		cur1 = next1
		cur2 = next2
	}

	if cur1 != nil {
        cur1.Next = nil
    }
}
