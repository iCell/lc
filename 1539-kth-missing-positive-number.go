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