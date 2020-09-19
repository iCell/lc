
func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }

    queue := []*Node{node}

    visited := make(map[*Node]*Node)
    visited[node] = &Node{node.Val, make([]*Node, 0, len(node.Neighbors))}

    for len(queue) > 0 {
        pop, temp := queue[0], queue[1:]
        queue = temp

        for _, neighbor := range pop.Neighbors {
            if _, ok := visited[neighbor]; !ok {
                queue = append(queue, neighbor)
                visited[neighbor] = &Node{neighbor.Val, make([]*Node, 0, len(neighbor.Neighbors))}
            }
            visited[pop].Neighbors = append(visited[pop].Neighbors, visited[neighbor])
        }
    }

    return visited[node]
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := make(map[*Node]*Node)
    dfs(node, visited)
    return visited[node]
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }

    if visited[node] != nil {
        return visited[node]
    }

    clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    visited[node] = clone

    for _, neighbor := range node.Neighbors {
        clone.Neighbors = append(
            clone.Neighbors,
            dfs(neighbor, visited),
        )
    }

    return clone
}
