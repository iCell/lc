func kthSmallest(root *TreeNode, k int) int {
    var idx int
    current, stack := root, []*TreeNode{root}
    for len(stack) != 0 || current != nil {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        top, temp := stack[len(stack)-1], stack[:len(stack)-1]
        stack = temp
        if idx == k-1 {
            return top.Val
        }
        current = top.Right
        idx++
    }
    return -1
}