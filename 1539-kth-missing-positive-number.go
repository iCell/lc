func findKthPositive(arr []int, k int) int {
	current, index, ans := 1, 0, -1
	for k > 0 {
		if current == arr[index] {
			if index + 1 < len(arr) {
				index += 1
			}
		} else {
			k -= 1
			ans = current
		}
		current += 1
	}
	return ans
}

// 比如 [2,3,4,7,11] 这个数组，如果考虑没有任何确实的数字，那么第一个数字应该为 1，也就是说
// arr[index] - index - 1 = 2 - 0 - 1 = 1，代表第 0 个位上之前缺失了 1 个字符，那么缺失的字符为 k + i

func findKthPositive(arr []int, k int) int {
	var index int
	for index < len(arr) {
		if arr[index] - index - 1 >= k {
			break
		}
		index++
	}
	return index + k
}

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] - mid - 1 >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left + k
}

func findKthPositive(arr []int, k int) int {
	memo, max := make(map[int]bool), math.MinInt32
	for _, num := range arr {
		memo[num] = true
		if num > max {
			max = num
		}
	}
	max += k

	var ans int
	for i := 1; i <= max; i++ {
		if memo[i] {
			continue
		}
		k--
		if k == 0 {
			ans = i
			break
		}
	}
	return ans
}