/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	visited := make(map[*Node]*Node)
	return dfs(head, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if visited[node] != nil {
		return visited[node]
	}
	
	clone := &Node{Val: node.Val}
	visited[node] = clone
	
	clone.Next = dfs(node.Next, visited)
	clone.Random = dfs(node.Random, visited)
	
	return clone
}