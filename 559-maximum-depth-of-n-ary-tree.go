
import "math"

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    if len(root.Children) == 0 {
        return 1
    }

    var heights []int
    for _, n := range root.Children {
        heights = append(heights, maxDepth(n))
    }

    depth := math.MinInt64
    for _, height := range heights {
        depth = max(depth, height)
    }

    return depth + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}