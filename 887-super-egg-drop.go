func superEggDrop(K int, N int) int {
	memo := make([][]int, K + 1)
	for i := range memo {
		memo[i] = make([]int, N + 1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return helper(K, N, memo)
}

func helper(k int, n int, memo [][]int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}
	if memo[k][n] != -1 {
		return memo[k][n]
	}
	
	ans := math.MaxInt32
	low, high := 1, n
	for low <= high {
		mid := (low + high) / 2
		broken := helper(k - 1, mid - 1, memo)
		notBroken := helper(k, n - mid, memo)
		if broken > notBroken {
			high = mid - 1
			ans = min(ans, broken + 1)
		} else {
			low = mid + 1
			ans = min(ans, notBroken + 1)
		}
	}
	
	memo[k][n] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

