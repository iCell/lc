func peakIndexInMountainArray(A []int) int {
    start, end := 0, len(A)-1
    for start < end {
        mid := (start + end) / 2
        if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
            return mid
        } else if A[mid] > A[mid-1] && A[mid] < A[mid+1] {
            start = mid + 1
        } else if A[mid] > A[mid+1] && A[mid] < A[mid-1] {
            end = mid - 1
        }
    }
    return start
}

func peakIndexInMountainArray(A []int) int {
    left, right := 0, len(A)-1
    for left < right {
        mid := (left + right) / 2
        if A[mid] < A[mid+1] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return right
}