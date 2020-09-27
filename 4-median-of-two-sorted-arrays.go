func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    left, right := (m+n+1)/2, (m+n+2)/2
    a := findKth(nums1, nums2, 0, 0, m-1, n-1, left)
    b := findKth(nums1, nums2, 0, 0, m-1, n-1, right)
    return float64(a+b) * 0.5
}

func findKth(nums1, nums2 []int, start1, start2 int, end1, end2 int, k int) int {
    len1, len2 := end1-start1+1, end2-start2+1
    if len1 == 0 {
        return nums2[start2+k-1]
    }
    if len2 == 0 {
        return nums1[start1+k-1]
    }
    if k == 1 {
        return min(nums1[start1], nums2[start2])
    }

    i := start1 + min(len1, k/2) - 1
    j := start2 + min(len2, k/2) - 1
    if nums1[i] >= nums2[j] {
        return findKth(nums1, nums2, start1, j+1, end1, end2, k-(j-start2+1))
    }
    return findKth(nums1, nums2, i+1, start2, end1, end2, k-(i-start1+1))
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

