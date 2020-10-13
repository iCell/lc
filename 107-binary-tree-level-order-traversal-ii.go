/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    var result [][]int

    for len(queue) != 0 {
        var level []int

        size := len(queue)
        for size > 0 {
            first, temp := queue[0], queue[1:]
            queue = temp

            level = append(level, first.Val)

            if first.Left != nil {
                queue = append(queue, first.Left)
            }
            if first.Right != nil {
                queue = append(queue, first.Right)
            }

            size--
        }

        result = append(result, level)
    }

    left, right := 0, len(result)-1
    for left < right {
        result[left], result[right] = result[right], result[left]
        left++
        right--
    }

    return result
}