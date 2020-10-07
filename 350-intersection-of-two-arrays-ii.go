import "sort"

func intersect(nums1 []int, nums2 []int) []int {
    memo := make(map[int]int, len(nums1))
    for _, num := range nums1 {
        memo[num] += 1
    }

    var result []int
    for _, num := range nums2 {
        _, exist := memo[num]
        if exist {
            memo[num] -= 1
            if memo[num] == 0 {
                delete(memo, num)
            }
            result = append(result, num)
        }
    }
    return result
}

func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    var result []int

    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            i++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            result = append(result, nums1[i])
            i++
            j++
        }
    }

    return result
}
