package main

import "math"

func jump(nums []int) int {
    var steps int

    start, end := 0, 1
    for end < len(nums) {
        maxJump := 0
        for i := start; i < end; i++ {
            maxJump = int(math.Max(float64(maxJump), float64(i+nums[i])))
        }
        start = end
        end = maxJump + 1

        steps++
    }

    return steps
}
