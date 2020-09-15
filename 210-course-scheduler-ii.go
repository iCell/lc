type Node struct {
    Num       int
    Neighbors []*Node
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    nodes := make([]*Node, 0, numCourses)
    degrees := make(map[int]int, numCourses)
    for i := 0; i < numCourses; i++ {
        degrees[i] = 0
        nodes = append(nodes, &Node{Num: i})
    }
    for _, prerequisite := range prerequisites {
        nodes[prerequisite[1]].Neighbors = append(
            nodes[prerequisite[1]].Neighbors,
            nodes[prerequisite[0]],
        )
        degrees[prerequisite[0]] += 1
    }

    queue := make([]*Node, 0)
    for n, v := range degrees {
        if v != 0 {
            continue
        }
        queue = append(queue, nodes[n])
    }

    var result []int
    for len(queue) != 0 {
        first, temp := queue[0], queue[1:]
        queue = temp
        result = append(result, first.Num)
        for _, neighbor := range first.Neighbors {
            degrees[neighbor.Num] -= 1
            if degrees[neighbor.Num] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    for _, v := range degrees {
        if v != 0 {
            return []int{}
        }
    }
    return result
}