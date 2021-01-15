/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func findRoot(tree []*Node) *Node {
	memo := make(map[int]bool)
	for _, node := range tree {
		for _, child := range node.Children {
			memo[child.Val] = true
		}
	}
	
	var root *Node
	for _, node := range tree {
		if !memo[node.Val] {
			root = node
			break
		}
	}
	return root
}