import "math"

func findClosestElements(arr []int, k int, x int) []int {
    i, j := 0, len(arr)-1
    for j-i+1 != k {
        left := math.Abs(float64(arr[i] - x))
        right := math.Abs(float64(arr[j] - x))
        if left == right {
            if arr[i] < arr[j] {
                j--
            } else {
                i++
            }
        } else if left > right {
            i++
        } else {
            j--
        }
    }
    return arr[i : j+1]
}