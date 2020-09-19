type Node struct {
    Num       int
    Neighbors map[*Node]int
}

func networkDelayTime(times [][]int, N int, K int) int {
    nodes := make(map[int]*Node, N)
    for i := 1; i <= N; i++ {
        nodes[i] = &Node{Num: i, Neighbors: make(map[*Node]int)}
    }
    for _, time := range times {
        source, destination, period := time[0], time[1], time[2]
        nodes[source].Neighbors[nodes[destination]] = period
    }

    queue, visited := []*Node{nodes[K]}, make(map[int]int)
    visited[K] = 0
    for len(queue) != 0 {
        first, temp := queue[0], queue[1:]
        queue = temp
        delete(nodes, first.Num)

        current := visited[first.Num]

        for neighbor, period := range first.Neighbors {
            used, exist := visited[neighbor.Num]
            if exist && used <= current+period {
                continue
            }
            queue = append(queue, neighbor)
            visited[neighbor.Num] = current + period
        }
    }
    if len(nodes) != 0 {
        return -1
    }

    var result int
    for _, v := range visited {
        if v > result {
            result = v
        }
    }
    return result
}