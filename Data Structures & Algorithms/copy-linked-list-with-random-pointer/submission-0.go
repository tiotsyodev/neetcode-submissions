/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
        return nil
    }

    resmap := map[*Node]*Node{}

	cur := head

	for cur != nil {
		created := &Node{
			Val: cur.Val,
		}
		resmap[cur] = created
		cur = cur.Next
	}

	cur = head

	for cur != nil {
		v, _ := resmap[cur]
		v.Next = resmap[cur.Next]
		v.Random = resmap[cur.Random]
		cur = cur.Next

	}

	return resmap[head]
}
