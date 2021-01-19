/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
	visited := make(map[*Node]bool)
	for p != nil {
		visited[p] = true
		p = p.Parent
	}
	
	for q != nil {
		if visited[q] {
			return q
		}
		q = q.Parent
	}
	
	return nil
}


func lowestCommonAncestor(p *Node, q *Node) *Node {
	pp, pq := p, q
	for pp != pq {
		pp, pq = pp.Parent, pq.Parent
		// 出现这种情况就代表 p 是 q 的 ancestor，或者反过来，此时就让两者强制相等就必定是对的
		if pp == nil {
			pp = q
		}
		if pq == nil {
			pq = p
		}
	}
	return pp
}
