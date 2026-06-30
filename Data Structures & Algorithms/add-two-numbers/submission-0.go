/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	dumcur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {

		sum := carry
		
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		dumcur.Next = &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10

		dumcur = dumcur.Next
		

	}

	return dummy.Next
}
