func numPairsDivisibleBy60(time []int) int {
	var ans int
	
	memo := make([]int, 60)
	for _, t := range time {
		memo[t % 60] += 1
	}
	
	ans += helper(memo[0]) + helper(memo[30])
	
	i, j := 1, 59
	for i < j {
		ans += memo[i] * memo[j]
		i++
		j--
	}

	return ans
}

func helper(n int) int {
	return n * (n - 1) / 2
}

