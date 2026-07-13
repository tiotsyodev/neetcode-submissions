/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	visited := map[*Node]*Node{}

	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if cl, ok := visited[node]; ok {
		return cl
	}

	clone := &Node{
		Val: node.Val,
		Neighbors: make([]*Node, 0),
	}

	visited[node] = clone

	for _, v := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(v, visited))
	}

	return clone
}
