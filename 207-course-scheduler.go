type Node struct {
    Num       int
    Neighbors []*Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    nodes := make([]*Node, 0, numCourses)
    degrees := make(map[int]int, numCourses)
    for i := 0; i < numCourses; i++ {
        nodes = append(nodes, &Node{Num: i})
        degrees[i] = 0
    }
    for _, prerequisite := range prerequisites {
        nodes[prerequisite[1]].Neighbors = append(nodes[prerequisite[1]].Neighbors, nodes[prerequisite[0]])
        degrees[prerequisite[0]] += 1
    }

    var queue []*Node
    for k, v := range degrees {
        if v == 0 {
            queue = append(queue, nodes[k])
        }
    }

    for len(queue) != 0 {
        first, temp := queue[0], queue[1:]
        queue = temp
        for _, neighbor := range first.Neighbors {
            degrees[neighbor.Num] -= 1
            if degrees[neighbor.Num] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    for _, v := range degrees {
        if v != 0 {
            return false
        }
    }
    return true
}