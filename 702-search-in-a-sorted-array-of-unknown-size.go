/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    if reader.get(0) == target {
        return 0
    }

    left, right := 0, 1
    for reader.get(right) < target {
        left = right
        right = right * 2
    }

    for left <= right {
        mid := (left + right) / 2
        num := reader.get(mid)
        if num == target {
            return mid
        } else if num > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}